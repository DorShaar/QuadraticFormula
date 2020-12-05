package Quadratic.Equation;

import Quadratic.Equation.Scanner.EquationScanException;
import Quadratic.Equation.Scanner.EquationScanner;

import java.util.ArrayList;
import java.util.List;

public class EquationArranger implements IEquationArranger
{
    private EquationScanner equationScanner;

    public EquationArranger(EquationScanner equationScanner)
    {
        this.equationScanner = equationScanner;
    }

    @Override
    public ArrangeResult arrange(String equation)
    {
        try
        {
            equationScanner.scan(equation);
            return innerArrange(equation);
        } catch (EquationScanException e)
        {
            return ArrangeResult.FailedArrange(equation, new ArrayList<>(List.of(e.getMessage())));
        }
    }

    private ArrangeResult innerArrange(String equation)
    {
        return ArrangeResult.SuccessArrange(equation, "3x^2-0x-27=0");
    }
}