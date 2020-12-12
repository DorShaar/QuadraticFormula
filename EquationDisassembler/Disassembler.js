import {EquationCoefficients} from "./EquationCoefficients.js"

// Get variable x and equation in the expected form of Ax+Bx+c=0.
// Extracts the coefficients A, b anc C.
export function disassamble(equation, variable) {
    var aResult = findA(equation, variable);
    aResult.secondDegreeIndex += 3; // + size of "x^2"

    var bResult = findB(equation, variable, aResult.secondDegreeIndex);
    bResult.firstDegreeIndex += 1; // + size of "x";

    const c = findC(equation, bResult.firstDegreeIndex);

    console.log("The coefficients of " + equation + " are: " + aResult.a + ", " + bResult.b + ", " + c);
    return new EquationCoefficients(equation, aResult.a, bResult.b, c);
}

function findA(equation, variable) {
    const secondDegreeVariable = variable + "^2";
    const secondDegreeIndex = equation.indexOf(secondDegreeVariable, 0);

    if (secondDegreeIndex == -1)
        throw new SyntaxError("Could not find second degree variable");

    console.log("Index of " + secondDegreeVariable + ": " + secondDegreeIndex);

    let a = equation.substring(0, secondDegreeIndex);
    if (a[0] == '+')
        a = a.substring(1);
    else if (a[0] != '-' && isNaN(a[0]))
        throw new SyntaxError("Expecting '+' or '-' signs only");

    console.log("A: " + a);

    return {
        a: a, 
        secondDegreeIndex: secondDegreeIndex,
    }
}

function findB(equation, variable, secondDegreeEndIndex) {
    const firstDegreeIndex = equation.indexOf(variable, secondDegreeEndIndex);

    if (firstDegreeIndex == -1)
        throw new SyntaxError("Could not find first degree variable");

    console.log("Index of " + variable + ": " + firstDegreeIndex);

    let b = equation.substring(secondDegreeEndIndex, firstDegreeIndex);
    if (b[0] == '+')
        b = b.substring(1);
    else if (b[0] != '-')
        throw new SyntaxError("Expecting '+' or '-' signs only");

    console.log("B: " + b);

    return {
        b: b, 
        firstDegreeIndex: firstDegreeIndex,
    }
}

function findC(equation, firstDegreeEndIndex) {
    const equalSignIndex = equation.indexOf("=", firstDegreeEndIndex);

    if (equalSignIndex == -1)
        throw new SyntaxError("Could not find equal sign");

    let c = equation.substring(firstDegreeEndIndex, equalSignIndex);
    if(c[0] == '+')
        c = c.substring(1);

    console.log("C: " + c);
    return c;
}

function createMessage() {
 // TODO use json
}