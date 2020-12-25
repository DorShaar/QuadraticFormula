using EquationSolver.Infra;
using Microsoft.Extensions.Logging.Abstractions;
using System.Threading.Tasks;
using Xunit;

namespace EquationSolverTests.Infra
{
    public class RedisCacheTests
    {
        [Theory]
        [InlineData()]
        public Task GetRootsIfExist_AsExpected()
        {
            RedisCache redisCache = new RedisCache(NullLogger<RedisCache>.Instance);
            return Task.CompletedTask;
        }

        [Fact]
        public Task SaveResult_AsExpected()
        {
            RedisCache redisCache = new RedisCache(NullLogger<RedisCache>.Instance);
            return Task.CompletedTask;
        }
    }
}