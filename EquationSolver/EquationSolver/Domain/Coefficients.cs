using System;

namespace EquationSolver.Domain
{
    public class Coefficients
    {
        public int A { get; private set; }
        public int B { get; private set; }
        public int C { get; private set; }

        public int OriginalA { get; }
        public int OriginalB { get; }
        public int OriginalC { get; }

        public Coefficients(int a, int b, int c)
        {
            OriginalA = a;
            OriginalB = b;
            OriginalC = c;

            MinimizeCoefficients();
        }

        public string GetCoefficientSignature()
        {
            return $"({A}, {B}, {C})";
        }

        private void MinimizeCoefficients()
        {
            int gcdAB = GCD(OriginalA, OriginalB);
            int gcdBC = GCD(OriginalA, OriginalC);

            int divideBy = Math.Min(gcdAB, gcdBC);

            if (divideBy == 1)
            {
                A = OriginalA;
                B = OriginalB;
                C = OriginalC;

                return;
            }

            A = OriginalA / divideBy;
            B = OriginalB / divideBy;
            C = OriginalC / divideBy;
        }

        private static int GCD(int a, int b)
        {
            if (a < 0)
                a = -a;

            if (b < 0)
                b = -b;

            while (a != 0 && b != 0)
            {
                if (a > b)
                    a %= b;
                else
                    b %= a;
            }

            return a == 0 ? b : a;
        }

        public override string ToString()
        {
            if (A == OriginalA)
                return $"({A}, {B}, {C})";

            return $"({OriginalA}, {OriginalB}, {OriginalC}) = ({A}, {B}, {C})";
        }
    }
}