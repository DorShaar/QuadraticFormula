package Quadratic.Equation.Scanner;

public enum ScanState
{
    DuringStart,
    DuringNumber,
    DuringVariable,
    DuringCoefficientSign,
    DuringExponentSign,
    DuringExponentNumber,
    DuringEqualSign,
    DuringMultiplySign,
}