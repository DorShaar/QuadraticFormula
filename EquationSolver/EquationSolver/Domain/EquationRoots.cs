using System.Diagnostics.CodeAnalysis;

namespace EquationSolver.Domain
{
    public class EquationRoots
    {
        public Coefficients Coefficients { get; }
        public double Root1 { get; }
        public double Root2 { get; }
        public bool HasResult { get; private set; } = true;

        public EquationRoots([NotNull] Coefficients coefficients, double root1, double root2)
        {
            if (coefficients.A == 0)
                HasResult = false;

            Coefficients = coefficients;
            Root1 = root1;
            Root2 = root2;
        }

        public static EquationRoots NoResult([NotNull] Coefficients coefficients)
        {
            return new EquationRoots(coefficients, 0, 0)
            {
                HasResult = false
            };
        }

        public override string ToString()
        {
            if (!HasResult)
                return "No Roots";

            if (Root1 == Root2)
                return $"({Root1})";

            return $"({Root1}, {Root2})";
        }
    }
}