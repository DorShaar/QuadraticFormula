package Scanner.Tests;

import Quadratic.Equation.Tests.Scanner.EquationScanException;
import Quadratic.Equation.Tests.Scanner.EquationScanner;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;
import org.junit.jupiter.params.provider.ValueSource;

import static org.junit.jupiter.api.Assertions.*;

public class EquationScannerTests
{
    @ParameterizedTest
    @CsvSource(
            {
                    "3x^2-0x-27=0,3,0,-27,0,0,0,x",
                    "3y^2-0y-27=0,3,0,-27,0,0,0,y",
                    "36x^2-0x+27=0,36,0,27,0,0,0,x",
                    "3*x^2-5*x-27=0,3,-5,-27,0,0,0,x",
                    "-3*x^2-0*x-27=4x^2 - 3x,-3,0,-27,4,-3,0,x",
                    "-3*x^2-50  *  x-27=4x^2 - 3x  + 2,-3,-50,-27,4,-3,2,x"
            })
    public void scan_ValidEquations_ScannedAsExpected(String input,
                                                       int secondDegreeLeft,
                                                       int firstDegreeLeft,
                                                       int freeNumberLeft,
                                                       int secondDegreeRight,
                                                       int firstDegreeRight,
                                                       int freeNumberRight,
                                                       char variable)
    {
        EquationScanner equationScanner = new EquationScanner();

        try
        {
            equationScanner.scan(input);
            assertEquals(secondDegreeLeft, equationScanner.secondDegreeCoefficientsOnLeft.get(0));
            assertEquals(firstDegreeLeft, equationScanner.firstDegreeCoefficientsOnLeft.get(0));
            assertEquals(freeNumberLeft, equationScanner.freeNumbersOnLeft.get(0));
            assertEquals(secondDegreeRight, equationScanner.secondDegreeCoefficientsOnRight.get(0));
            assertEquals(firstDegreeRight, equationScanner.firstDegreeCoefficientsOnRight.get(0));
            assertEquals(freeNumberRight, equationScanner.freeNumbersOnRight.get(0));
            assertEquals(variable, equationScanner.variableSign);
        } catch (EquationScanException e)
        {
            fail();
        }
    }

    @ParameterizedTest
    @ValueSource(strings =
            {
                    "3xx^0-0x-27=0",
                    "3xx^0-0x-27==0",
                    "3xx^3-0x-27=0",
                    "3xx^22-0x-27=0",
                    "3x^2--0x-27=0",
                    "3x^2-0**x-27=0",
                    "3x^2-0*x-*27=0",
                    "3x^2-0*x-27=",
                    "3x^2-0y-27=0",
                    "abc"
            })
    public void scan_InvalidEquations_ThrowsEquationScanException(String input)
    {
        EquationScanner equationScanner = new EquationScanner();

        assertThrows(EquationScanException.class, () -> equationScanner.scan(input));
    }

    @Test
    public void scan_ValidSpecificEquation1_ScannedAsExpected()
    {
        EquationScanner equationScanner = new EquationScanner();

        String equation = "3x^1-0x-27=0";
        try
        {
            equationScanner.scan(equation);
            assertEquals(0, equationScanner.secondDegreeCoefficientsOnLeft.get(0));
            assertEquals(3, equationScanner.firstDegreeCoefficientsOnLeft.get(0));
            assertEquals(0, equationScanner.firstDegreeCoefficientsOnLeft.get(1));
            assertEquals(-27, equationScanner.freeNumbersOnLeft.get(0));
            assertEquals(0, equationScanner.secondDegreeCoefficientsOnRight.get(0));
            assertEquals(0, equationScanner.firstDegreeCoefficientsOnRight.get(0));
            assertEquals(0, equationScanner.freeNumbersOnRight.get(0));
            assertEquals('x', equationScanner.variableSign);
        } catch (EquationScanException e)
        {
            fail();
        }
    }

    @Test
    public void scan_ValidSpecificEquation2_ScannedAsExpected()
    {
        EquationScanner equationScanner = new EquationScanner();

        String equation = "3x^0-0x-27=0";
        try
        {
            equationScanner.scan(equation);
            assertEquals(0, equationScanner.secondDegreeCoefficientsOnLeft.get(0));
            assertEquals(0, equationScanner.firstDegreeCoefficientsOnLeft.get(0));
            assertEquals(3, equationScanner.freeNumbersOnLeft.get(0));
            assertEquals(-27, equationScanner.freeNumbersOnLeft.get(1));
            assertEquals(0, equationScanner.secondDegreeCoefficientsOnRight.get(0));
            assertEquals(0, equationScanner.firstDegreeCoefficientsOnRight.get(0));
            assertEquals(0, equationScanner.freeNumbersOnRight.get(0));
            assertEquals('x', equationScanner.variableSign);
        } catch (EquationScanException e)
        {
            fail();
        }
    }
}