using EquationSolver.Domain;
using EquationSolver.Infra;
using Microsoft.Extensions.Logging;
using System.Diagnostics.CodeAnalysis;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;

namespace EquationSolver.App
{
    public class EquationManager
    {
        private readonly Solver mSolver;
        private readonly RedisCache mCache;
        private readonly ILogger<EquationManager> mLogger;

        public EquationManager(
            [NotNull] Solver solver,
            [NotNull] RedisCache cache,
            [NotNull] ILogger<EquationManager> logger)
        {
            mSolver = solver;
            mCache = cache;
            mLogger = logger;
        }

        public async Task<EquationRoots> HandleEquation(Coefficients coefficients, CancellationToken cancellationToken)
        {
            EquationRoots equationRoots =
                await mCache.GetRootsIfExist(coefficients.GetCoefficientSignature(), cancellationToken).ConfigureAwait(false);

            if (equationRoots != null)
                return equationRoots;

            equationRoots = mSolver.Solve(coefficients);
            await mCache.SaveResult(coefficients.GetCoefficientSignature(), equationRoots, cancellationToken: cancellationToken)
                .ConfigureAwait(false);

            return equationRoots;
        }

        public Task SendMessage(EquationRoots equationRootsMessage)
        {
            string serializedString = JsonSerializer.Serialize(equationRootsMessage);
            mLogger.LogInformation($"Sending message: {serializedString}");

            return Task.CompletedTask;
        }
    }
}