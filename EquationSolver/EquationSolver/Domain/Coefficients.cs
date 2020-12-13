namespace EquationSolver.Domain
{
    public class Coefficients
    {
        public int A { get; }
        public int B { get; }
        public int C { get; }

        public Coefficients(int a, int b, int c)
        {
            A = a;
            B = b;
            C = c;
        }
    }
}