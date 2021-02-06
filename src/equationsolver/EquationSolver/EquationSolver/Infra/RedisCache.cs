using EquationSolver.Domain;
using Microsoft.Extensions.Logging;
using Polly;
using StackExchange.Redis;
using System;
using System.Diagnostics.CodeAnalysis;
using System.Text;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;

namespace EquationSolver.Infra
{
    public class RedisCache : IDisposable, IAsyncDisposable
    {
        private bool mIsDisposed;

        private readonly TimeSpan mTTL = TimeSpan.FromMinutes(3);
        private readonly int mMaxRetries = 3;

        private readonly ILogger<RedisCache> mlogger;
        private readonly ConnectionMultiplexer mMultiplexer = ConnectionMultiplexer.Connect("localhost,allowAdmin=true");

        public RedisCache([NotNull] ILogger<RedisCache> logger)
        {
            mlogger = logger;
        }

        public async Task<EquationRoots> GetRootsIfExist(string key, CancellationToken cancellationToken)
        {
            mlogger.LogDebug($"Searching if there is an existing result for {key}");
            IDatabase db = mMultiplexer.GetDatabase();

            byte[] deserializedEquationRoots = await OperateWithRetry(() => db.StringGetAsync(key)).ConfigureAwait(false);

            if (deserializedEquationRoots == null)
            {
                mlogger.LogDebug($"No roots for {key} in cache");
                return null;
            }

            EquationRoots equationRoots = Deserialize(deserializedEquationRoots);
            mlogger.LogDebug($"Found exiting roots for {key} in cache: {equationRoots}");
            return equationRoots;
        }

        public async Task SaveResult(string key, EquationRoots equationRoots,
            TimeSpan timeSpan = default, CancellationToken cancellationToken = default)
        {
            mlogger.LogDebug($"Saving value {equationRoots} with key {key}");
            IDatabase db = mMultiplexer.GetDatabase();

            byte[] serializedEquationRoots = Serialize(equationRoots);

            TimeSpan ttl = timeSpan == default ? mTTL : timeSpan;

            bool wasSet = await OperateWithRetry(() => db.StringSetAsync(key, serializedEquationRoots, ttl)).ConfigureAwait(false);

            if (!wasSet)
                mlogger.LogError($"Could not set value {equationRoots} with key {equationRoots}");
        }

        private async Task<T> OperateWithRetry<T>(Func<Task<T>> redisOperation)
        {
            int retryAttempt = 1;

            return await Policy.Handle<Exception>()
                .WaitAndRetry(mMaxRetries, retryAttempt => TimeSpan.FromMilliseconds(200 * retryAttempt),
                    (_, _) =>
                    {
                        mlogger.LogWarning($"Could not execute redis operation, attempt: {retryAttempt} out of {mMaxRetries}");
                        ++retryAttempt;
                    }
                )
                .Execute(async () => await redisOperation().ConfigureAwait(false))
                .ConfigureAwait(false);
        }

        public void Flush(string dbName)
        {
            IServer server = mMultiplexer.GetServer(dbName);
            server.FlushDatabase();
        }

        private static byte[] Serialize(EquationRoots objectToSerialize)
        {
            if (objectToSerialize == null)
                return null;

            string serializedString = JsonSerializer.Serialize(objectToSerialize);
            return Encoding.ASCII.GetBytes(serializedString);
        }

        private static EquationRoots Deserialize(byte[] byteArray)
        {
            string deserializedString = Encoding.ASCII.GetString(byteArray);
            return JsonSerializer.Deserialize<EquationRoots>(deserializedString);
        }

        public void Dispose()
        {
            Dispose(true);
            GC.SuppressFinalize(this);
        }

        protected virtual void Dispose(bool disposing)
        {
            if (mIsDisposed) return;

            if (disposing)
            {
                mMultiplexer.Dispose();
            }

            mIsDisposed = true;
        }

        public ValueTask DisposeAsync()
        {
            Dispose();
            return default;
        }
    }
}