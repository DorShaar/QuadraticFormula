package Quadratic.Equation.Tests.Scanner;

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