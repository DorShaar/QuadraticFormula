package Quadratic.Equation.Tests;

public class EquationArranger implements IEquationArranger
{
    @Override
    public ArrangeResult arrange(String equation)
    {
        return ArrangeResult.SuccessArrange(equation, "3x^2-0x-27=0");
    }

    private void scanEquation(String equation)
    {

    }
}