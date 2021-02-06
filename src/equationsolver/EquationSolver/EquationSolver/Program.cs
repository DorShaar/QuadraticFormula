using EquationSolver.Infra;
using Microsoft.Extensions.Logging.Abstractions;
using System;
using System.Threading.Tasks;

namespace EquationSolver
{
    public static class Program
    {
        public static async Task Main()
        {
            Console.WriteLine("Hello World!");

            using RedisCache redisCache = new RedisCache(NullLogger<RedisCache>.Instance);
            //await redisCache.Set().ConfigureAwait(false);
            //await redisCache.Get().ConfigureAwait(false);
        }
    }
}