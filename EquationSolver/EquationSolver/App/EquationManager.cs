using EquationSolver.Domain;
using EquationSolver.Infra;
using Microsoft.Extensions.Logging;
using System.Diagnostics.CodeAnalysis;
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

        public async Task HandleEquation(Coefficients coefficients, CancellationToken cancellationToken)
        {
            EquationRoots equationRoots =
                await mCache.GetRootsIfExist(coefficients.GetCoefficientSignature(), cancellationToken).ConfigureAwait(false);

            if (equationRoots == null)
            {
                equationRoots = mSolver.Solve(coefficients);
                await mCache.SaveResult(coefficients.GetCoefficientSignature(), equationRoots, cancellationToken)
                    .ConfigureAwait(false);

                return;
            }
        }
    }
}