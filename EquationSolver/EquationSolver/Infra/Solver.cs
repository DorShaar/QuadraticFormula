using EquationSolver.Domain;
using Microsoft.Extensions.Logging;
using System;
using System.Diagnostics.CodeAnalysis;

namespace EquationSolver.Infra
{
    public class Solver
    {
        private readonly ILogger<Solver> mLogger;

        public Solver([NotNull] ILogger<Solver> logger)
        {
            mLogger = logger;
        }

        public EquationRoots Solve(Coefficients coefficients)
        {
            int a = coefficients.A;
            int b = coefficients.B;
            int c = coefficients.C;

            mLogger.LogDebug($"Calcualting roots for coefficient a: {a}, b: {b} and c: {c}");

            if (a == 0)
            {
                mLogger.LogInformation($"Coefficients a: {a}, b: {b} and c: {c} has no quadratic root since a = 0");
                return EquationRoots.NoResult(coefficients);
            }

            double delta = Math.Pow(b, 2) - (4 * a * c);

            if (delta < 0)
            {
                mLogger.LogInformation($"Coefficients a: {a}, b: {b} and c: {c} has no quadratic roots. Delta = {delta}");
                return EquationRoots.NoResult(coefficients);
            }

            double x1 = (-b + Math.Sqrt(delta)) / 2 * a;
            double x2 = (-b + Math.Sqrt(delta)) / 2 * a;

            return new EquationRoots(coefficients, x1, x2);
        }
    }
}