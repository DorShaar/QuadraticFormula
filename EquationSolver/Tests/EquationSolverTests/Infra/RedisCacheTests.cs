using EquationSolver.Domain;
using EquationSolver.Infra;
using Microsoft.Extensions.Logging.Abstractions;
using System;
using System.Threading;
using System.Threading.Tasks;
using Xunit;

namespace EquationSolverTests.Infra
{
    public class RedisCacheTests : IDisposable
    {
        private readonly RedisCache mRedisCache = new RedisCache(NullLogger<RedisCache>.Instance);

        [Fact]
        public async Task SaveResult_And_GetRootsIfExist_AsExpected()
        {
            Coefficients coefficients = new Coefficients(1, 2, 3);

            EquationRoots equationRoots =
                await mRedisCache.GetRootsIfExist(coefficients.GetCoefficientSignature(), CancellationToken.None).ConfigureAwait(false);

            Assert.Null(equationRoots);

            const double root1 = 1;
            const double root2 = 5.5;

            equationRoots = new EquationRoots(coefficients, root1, root2);

            await mRedisCache.SaveResult(
                coefficients.GetCoefficientSignature(), equationRoots, TimeSpan.FromSeconds(2), CancellationToken.None)
               .ConfigureAwait(false);

            equationRoots =
                await mRedisCache.GetRootsIfExist(coefficients.GetCoefficientSignature(), CancellationToken.None).ConfigureAwait(false);

            Assert.Equal(root1, equationRoots.Root1);
            Assert.Equal(root2, equationRoots.Root2);
        }

        public void Dispose()
        {
            mRedisCache.Dispose();
        }
    }
}