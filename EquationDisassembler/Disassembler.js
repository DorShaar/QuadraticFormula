import {EquationCoefficients} from "./EquationCoefficients.js"

// Get variable x and equation in the expected form of Ax+Bx+c=0.
// Extracts the coefficients A, b anc C.
export function disassamble(equation, variable) {
    let a, b, c;

    var aResult = findA(equation, variable);
    aResult.secondDegreeIndex += 3; // + size of "x^2"
    a = aResult.a;

    var bResult = findB(equation, variable, aResult.secondDegreeIndex);
    bResult.firstDegreeIndex += 1; // + size of "x";
    b = bResult.b;

    c = findC(equation, bResult.firstDegreeIndex);

    console.log("The coefficients of " + equation + " are: " + a + ", " + b + ", " + c);
    return new EquationCoefficients(equation, a, b, c);
}

function findA(equation, variable) {
    const secondDegreeVariable = variable + "^2";
    const secondDegreeIndex = equation.indexOf(secondDegreeVariable, 0);
    console.log("Index of " + secondDegreeVariable + ": " + secondDegreeIndex);

    let a = equation.substring(0, secondDegreeIndex);
    if(a[0] == '+')
        a = a.substring(1);

    console.log("A: " + a);

    return {
        a: a, 
        secondDegreeIndex: secondDegreeIndex,
    }
}

function findB(equation, variable, secondDegreeEndIndex) {
    const firstDegreeIndex = equation.indexOf(variable, secondDegreeEndIndex);
    console.log("Index of " + variable + ": " + firstDegreeIndex);

    let b = equation.substring(secondDegreeEndIndex, firstDegreeIndex);
    if(b[0] == '+')
        b = b.substring(1);

    console.log("B: " + b);

    return {
        b: b, 
        firstDegreeIndex: firstDegreeIndex,
    }
}

function findC(equation, firstDegreeEndIndex) {
    const equalSignIndex = equation.indexOf("=", firstDegreeEndIndex);
    let c = equation.substring(firstDegreeEndIndex, equalSignIndex);
    if(c[0] == '+')
        c = c.substring(1);

    console.log("C: " + c);
    return c;
}

function createMessage() {
 // TODO use json
}