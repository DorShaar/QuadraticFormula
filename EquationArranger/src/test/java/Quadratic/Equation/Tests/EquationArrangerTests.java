package Quadratic.Equation.Tests;

import Quadratic.Equation.ArrangeResult;
import Quadratic.Equation.EquationArranger;
import Quadratic.Equation.Scanner.EquationScanner;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;
import org.junit.jupiter.params.provider.ValueSource;

import static org.junit.jupiter.api.Assertions.*;

public class EquationArrangerTests
{
    @ParameterizedTest
    @CsvSource(
            {
                "4x^2 - 22 =x^2+5,3x^2-0x-27=0",
                "4x^2- 22,4x^2-22=0",
                "-5x + 3,0x^2-5x+3=0",
                "y ^ 2 + y  = 1,x^2+x-1=0,true",
            })
    public void arrange_ValidEquations_ArrangedAsExpected(String input,
                                                          String expected)
    {
        EquationScanner equationScanner = new EquationScanner();
        EquationArranger equationArranger = new EquationArranger(equationScanner);

        ArrangeResult arrangeResult = equationArranger.arrange(input);
        assertTrue(arrangeResult.isArrangeSucceeded());
        assertEquals(expected, arrangeResult.arrangedEquation());
    }

    @ParameterizedTest
    @ValueSource(strings = {"abc", "y ^ 2 + x  = 1", "x^3 + x^2 + 1 = 1"})
    public void arrange_InvalidEquations_ArrangedAsExpected(String input)
    {
        EquationScanner equationScanner = new EquationScanner();
        EquationArranger equationArranger = new EquationArranger(equationScanner);

        ArrangeResult arrangeResult = equationArranger.arrange(input);
        assertFalse(arrangeResult.isArrangeSucceeded());
    }
}