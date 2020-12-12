package Quadratic.Equation.Scanner;

import java.util.ArrayList;
import java.util.List;

public class EquationScanner
{
    public List<Integer> secondDegreeCoefficientsOnLeft;
    public List<Integer> firstDegreeCoefficientsOnLeft;
    public List<Integer> freeNumbersOnLeft;
    public List<Integer> secondDegreeCoefficientsOnRight;
    public List<Integer> firstDegreeCoefficientsOnRight;
    public List<Integer> freeNumbersOnRight;
    public char variableSign;

    private Boolean isAfterEqualSign;
    private ScanState scanState;
    private List<Character> charactersGroup;
    private char coefficientSign;

    public void scan(String equation) throws EquationScanException
    {
        init();
        collectCoefficients(equation);
        end();
    }

    private void init()
    {
        secondDegreeCoefficientsOnLeft = new ArrayList<>(1);
        firstDegreeCoefficientsOnLeft = new ArrayList<>(1);
        freeNumbersOnLeft = new ArrayList<>(1);
        secondDegreeCoefficientsOnRight = new ArrayList<>(1);
        firstDegreeCoefficientsOnRight = new ArrayList<>(1);
        freeNumbersOnRight = new ArrayList<>(1);
        variableSign = ' ';

        isAfterEqualSign = false;
        scanState = ScanState.DuringStart;
        charactersGroup = new ArrayList<>();
        coefficientSign = '+';
    }

    private void end()
    {
        if (secondDegreeCoefficientsOnLeft.size() == 0)
            secondDegreeCoefficientsOnLeft.add(0);

        if (firstDegreeCoefficientsOnLeft.size() == 0)
            firstDegreeCoefficientsOnLeft.add(0);

        if (freeNumbersOnLeft.size() == 0)
            freeNumbersOnLeft.add(0);

        if (secondDegreeCoefficientsOnRight.size() == 0)
            secondDegreeCoefficientsOnRight.add(0);

        if (firstDegreeCoefficientsOnRight.size() == 0)
            firstDegreeCoefficientsOnRight.add(0);

        if (freeNumbersOnRight.size() == 0)
            freeNumbersOnRight.add(0);
    }

    private void collectCoefficients(String equation) throws EquationScanException
    {
        for (char ch : equation.toCharArray())
        {
            if (Character.isWhitespace(ch))
                continue;

            validateCharacterWithState(ch);

            if (Character.isDigit(ch))
            {
                handleDigit(ch);
                continue;
            }

            if (Character.isAlphabetic(ch))
            {
                handleVariable(ch);
                continue;
            }

            handleSignCharacter(ch);
        }

        handleScanEnd();
    }

    private void handleDigit(char ch) throws EquationScanException
    {
        if (scanState == ScanState.DuringExponentSign)
        {
            handleEndOfExponentScan(ch);
            scanState = ScanState.DuringExponentNumber;
            return;
        }

        scanState = ScanState.DuringNumber;
        charactersGroup.add(ch);
    }

    private void handleVariable(char ch) throws EquationScanException
    {
        if (scanState == ScanState.DuringCoefficientSign ||
            scanState == ScanState.DuringEqualSign ||
            scanState == ScanState.DuringStart)
        {
            charactersGroup.add('1');
        }

        scanState = ScanState.DuringVariable;

        if (variableSign == ' ')
        {
            variableSign = ch;
            return;
        }

        if (variableSign == ch)
            return;

        throw new EquationScanException("Found two different variables in the equation. " +
                "Variables: " + variableSign + ", " + ch);
    }

    private void handleSignCharacter(char ch) throws EquationScanException
    {
        switch (ch)
        {
            case '^':
                handleExponentSign();
                break;

            case '-':
            case '+':
                handleCoefficientSign(ch);
                break;

            case '=':
                handleEqualSign(ch);
                break;

            case '*':
                handleMultiplySign(ch);
                break;

            default:
                throw new EquationScanException("Unexpected character '" + ch + "'");
        }
    }

    private void handleExponentSign()
    {
        scanState = ScanState.DuringExponentSign;
    }

    private void handleCoefficientSign(char ch) throws EquationScanException
    {
        if (scanState == ScanState.DuringNumber)
            addFreeNumber();

        if (scanState == ScanState.DuringVariable)
            addCoefficient();

        scanState = ScanState.DuringCoefficientSign;
        coefficientSign = ch;
    }

    private void handleEqualSign(char ch) throws EquationScanException
    {
        if (isAfterEqualSign)
            throw new EquationScanException("Equation has two '=' signs");

        if (scanState == ScanState.DuringVariable)
            addCoefficient();

        if (scanState == ScanState.DuringNumber)
            addFreeNumber();

        scanState = ScanState.DuringEqualSign;
        isAfterEqualSign = true;
    }

    private void handleMultiplySign(char ch)
    {
        scanState = ScanState.DuringMultiplySign;
    }

    private void validateCharacterWithState(char ch) throws EquationScanException
    {
        switch (scanState)
        {
            case DuringStart:
            case DuringEqualSign:
            {
                if (Character.isDigit(ch) || Character.isAlphabetic(ch) || ch == '-' || ch == '+')
                    return;

                break;
            }

            case DuringNumber:
            {
                if (Character.isDigit(ch) || Character.isAlphabetic(ch) ||
                        ch == '-' || ch == '+' || ch == '*' || ch == '=')
                    return;

                break;
            }

            case DuringCoefficientSign:
            {
                if (Character.isDigit(ch) || Character.isAlphabetic(ch))
                    return;

                break;
            }

            case DuringMultiplySign:
            {
                if (Character.isAlphabetic(ch))
                    return;

                break;
            }

            case DuringVariable:
            {
                if (ch == '-' || ch == '+' || ch == '=' || ch == '^')
                    return;

                break;
            }

            case DuringExponentSign:
            {
                if (Character.isDigit(ch))
                    return;

                break;
            }

            case DuringExponentNumber:
            {
                if (ch == '-' || ch == '+')
                    return;

                break;
            }

            default:
                throw new EquationScanException("Unexpected case encountered " + scanState);
        }

        throw new EquationScanException(
                "Invalid position for character '" + ch + "'. Current State: " + scanState);
    }

    private void handleScanEnd() throws EquationScanException
    {
        switch (scanState)
        {
            case DuringExponentNumber:
                return;

            case DuringNumber:
                addFreeNumber();
                break;

            case DuringVariable:
                addCoefficient();
                break;

            default:
                throw new EquationScanException("Invalid end of equation");
        }
    }

    private void handleEndOfExponentScan(char ch) throws EquationScanException
    {
        try
        {
            if (ch == '2')
            {
                addCoefficientOfExponent2();
                return;
            }

            if (ch == '1')
            {
                addCoefficient();
                return;
            }

            if (ch == '0')
            {
                addFreeNumber();
                return;
            }

            throw new EquationScanException("Found invalid exponent character " + ch);
        }
        finally
        {
            scanState = ScanState.DuringExponentNumber;
        }
    }

    private void addCoefficientOfExponent2() throws EquationScanException
    {
        int number = CreateNumberFromCharactersGroupAndCoefficientSign();
        if (isAfterEqualSign)
        {
            secondDegreeCoefficientsOnRight.add(number);
            return;
        }

        secondDegreeCoefficientsOnLeft.add(number);
    }

    private void addCoefficient() throws EquationScanException
    {
        int number = CreateNumberFromCharactersGroupAndCoefficientSign();
        if (isAfterEqualSign)
        {
            firstDegreeCoefficientsOnRight.add(number);
            return;
        }

        firstDegreeCoefficientsOnLeft.add(number);
    }

    private void addFreeNumber() throws EquationScanException
    {
        int number = CreateNumberFromCharactersGroupAndCoefficientSign();
        if (isAfterEqualSign)
        {
            freeNumbersOnRight.add(number);
            return;
        }

        freeNumbersOnLeft.add(number);
    }

    private int CreateNumberFromCharactersGroupAndCoefficientSign() throws EquationScanException
    {
        int number = CreateNumberFromCharactersGroup();

        try
        {
            if (coefficientSign == '-')
                return -number;

            if (coefficientSign == '+')
                return number;

            throw new EquationScanException("Unexpected coefficient sign " + coefficientSign);
        }
        finally
        {
            coefficientSign = '+';
        }
    }

    private int CreateNumberFromCharactersGroup()
    {
        int number = 0;
        for(int i = 0; i < charactersGroup.size(); ++i)
        {
            number = number * 10;
            number += Character.getNumericValue(charactersGroup.get(i));
        }

        charactersGroup.clear();
        return number;
    }
}