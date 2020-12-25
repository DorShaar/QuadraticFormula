using EquationSolver.Domain;
using Xunit;

namespace EquationSolverTests.Domain
{
    public class CoefficientsTests
    {
        [Theory]
        [InlineData(4, 6, 8, 2, 3, 4, "(4, 6, 8) = (2, 3, 4)")]
        [InlineData(-4, 6, 8, -2, 3, 4, "(-4, 6, 8) = (-2, 3, 4)")]
        [InlineData(4, -6, 8, 2, -3, 4, "(4, -6, 8) = (2, -3, 4)")]
        [InlineData(4, 6, -8, 2, 3, -4, "(4, 6, -8) = (2, 3, -4)")]
        [InlineData(-4, 6, -8, -2, 3, -4, "(-4, 6, -8) = (-2, 3, -4)")]
        [InlineData(-4, -6, 8, -2, -3, 4, "(-4, -6, 8) = (-2, -3, 4)")]
        [InlineData(4, -6, -8, 2, -3, -4, "(4, -6, -8) = (2, -3, -4)")]
        [InlineData(-4, -6, -8, -2, -3, -4, "(-4, -6, -8) = (-2, -3, -4)")]
        [InlineData(4, 8, 12, 1, 2, 3, "(4, 8, 12) = (1, 2, 3)")]
        [InlineData(5, 7, 12, 5, 7, 12, "(5, 7, 12)")]
        [InlineData(5, 15, 12, 5, 15, 12, "(5, 15, 12)")]
        public void Ctor_AndToString_AsExpected(int a, int b, int c, int minimizedA, int minimizedB, int minimizedC,
            string expectedToString)
        {
            Coefficients coefficients = new Coefficients(a, b, c);

            Assert.Equal(minimizedA, coefficients.A);
            Assert.Equal(minimizedB, coefficients.B);
            Assert.Equal(minimizedC, coefficients.C);
            Assert.Equal(expectedToString, coefficients.ToString());
        }
    }
}