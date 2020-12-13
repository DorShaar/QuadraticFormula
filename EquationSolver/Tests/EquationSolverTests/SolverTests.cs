using EquationSolver.Domain;
using EquationSolver.Infra;
using Microsoft.Extensions.Logging.Abstractions;
using Xunit;

namespace EquationSolverTests
{
    public class SolverTests
    {
        [Theory]
        [InlineData(1, -5, 6, 3, -3)]
        public void Solve_EquationWithRoots_RootsAreCorrect(int a, int b, int c, double x1, double x2)
        {
            Solver solver = new Solver(NullLogger<Solver>.Instance);
            EquationRoots equationRoots = solver.Solve(a, b, c);

            Assert.True(equationRoots.HasResult);
            Assert.Equal(x1, equationRoots.Root1);
            Assert.Equal(x2, equationRoots.Root2);
        }

        [Fact]
        public void Solve_EquationWithOneRoot_RootIsCorrect()
        {
            Assert.True(false);
        }

        [Fact]
        public void Solve_EquationWithoutRoots_ReturnsNoResult()
        {
            Assert.True(false);
        }

        [Fact]
        public void Solve_EquationWithZeroACoefficientOneRoot_ReturnsNoResult()
        {
            Assert.True(false);
        }
    }
}