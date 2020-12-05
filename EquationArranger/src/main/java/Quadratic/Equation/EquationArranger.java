package Quadratic.Equation;

import Quadratic.Equation.Scanner.EquationScanException;
import Quadratic.Equation.Scanner.EquationScanner;
import org.apache.log4j.Logger;

import java.util.ArrayList;
import java.util.List;

public class EquationArranger implements IEquationArranger
{
    private EquationScanner equationScanner;
    final static Logger logger = Logger.getLogger(EquationArranger.class);

    public EquationArranger(EquationScanner equationScanner)
    {
        this.equationScanner = equationScanner;
    }

    @Override
    public ArrangeResult arrange(String equation)
    {
        logger.info("Arranging equation " + equation);

        try
        {
            equationScanner.scan(equation);
            return innerArrange(equation);
        } catch (EquationScanException ex)
        {
            logger.error("Arranging equation " + equation + " failed", ex);
            return ArrangeResult.FailedArrange(equation, new ArrayList<>(List.of(ex.getMessage())));
        }
    }

    private ArrangeResult innerArrange(String equation)
    {
        return ArrangeResult.SuccessArrange(equation, "3x^2-0x-27=0");
    }
}