using EquationSolver.Domain;
using EquationSolver.Infra;
using Microsoft.Extensions.Logging.Abstractions;
using Xunit;

namespace EquationSolverTests.Infra
{
    public class SolverTests
    {
        [Theory]
        [InlineData(1, -5, 6, 3, 2)]
        [InlineData(2, 4, -6, 1, -3)]
        public void Solve_EquationWithRoots_RootsAreCorrect(int a, int b, int c, double x1, double x2)
        {
            Solver solver = new Solver(NullLogger<Solver>.Instance);
            EquationRoots equationRoots = solver.Solve(new Coefficients(a, b, c));

            Assert.True(equationRoots.HasResult);
            Assert.Equal(x1, equationRoots.Root1);
            Assert.Equal(x2, equationRoots.Root2);
        }

        [Theory]
        [InlineData(-4, 12, -9, 1.5)]
        public void Solve_EquationWithOneRoot_RootIsCorrect(int a, int b, int c, double x)
        {
            Solver solver = new Solver(NullLogger<Solver>.Instance);
            EquationRoots equationRoots = solver.Solve(new Coefficients(a, b, c));

            Assert.True(equationRoots.HasResult);
            Assert.Equal(x, equationRoots.Root1);
        }

        [Theory]
        [InlineData(1, -3, 4)]
        public void Solve_EquationWithoutRoots_ReturnsNoResult(int a, int b, int c)
        {
            Solver solver = new Solver(NullLogger<Solver>.Instance);
            EquationRoots equationRoots = solver.Solve(new Coefficients(a, b, c));

            Assert.False(equationRoots.HasResult);
        }

        [Theory]
        [InlineData(0, -3, 4)]
        public void Solve_EquationWithZeroACoefficientOneRoot_ReturnsNoResult(int a, int b, int c)
        {
            Solver solver = new Solver(NullLogger<Solver>.Instance);
            EquationRoots equationRoots = solver.Solve(new Coefficients(a, b, c));

            Assert.False(equationRoots.HasResult);
        }
    }
}