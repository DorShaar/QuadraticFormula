using EquationSolver.Domain;
using Microsoft.Extensions.Logging;
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

            byte[] deserializedEquationRoots = await db.StringGetAsync(key).ConfigureAwait(false);

            if (deserializedEquationRoots == null)
            {
                mlogger.LogDebug($"No roots for {key} in cache");
                return null;
            }

            EquationRoots equationRoots = Deserialize(deserializedEquationRoots);
            mlogger.LogDebug($"Found exiting roots for {key} in cache: {equationRoots}");
            return equationRoots;
        }

        public async Task SaveResult(string key, EquationRoots equationRoots, CancellationToken cancellationToken)
        {
            mlogger.LogDebug($"Saving value {equationRoots} with key {key}");
            IDatabase db = mMultiplexer.GetDatabase();

            byte[] serializedEquationRoots = Serialize(equationRoots);

            // TODO polly
            bool wasSet = await db.StringSetAsync(key, serializedEquationRoots, TimeSpan.FromSeconds(60)).ConfigureAwait(false);

            if (!wasSet)
                mlogger.LogError($"Could not set value {equationRoots} with key {equationRoots}");
        }

        public void Flush(string dbName)
        {
            IServer server = mMultiplexer.GetServer(dbName);
            server.FlushDatabase();
        }

        // TODO use + write using polly.
        //private static void HandleRedisConnectionError(RedisConnectionException ex)
        //{
        //    if (attemptingToConnect) return;
        //    try
        //    {
        //        Policy
        //            .Handle<Exception>()
        //            .WaitAndRetry(3, retryAttempt => TimeSpan.FromSeconds(Math.Pow(2, retryAttempt)),
        //                (exception, timeSpan) =>
        //                {
        //                    Debug.WriteLine("Redis retry attempt" + exception.Message);
        //                }
        //            )
        //            .Execute(() =>
        //            {
        //                attemptingToConnect = true;
        //                RedisConnection.Reconnect();
        //            });
        //    }
        //    catch (Exception)
        //    {
        //        throw;
        //    }
        //    finally
        //    {
        //        attemptingToConnect = false;
        //    }
        //}

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