import {DisassembledEquationMessage} from "./DisassembledEquationMessage.js"
import winston from 'winston';
import {format} from 'logform';

const logger = winston.createLogger({
  level: 'debug',
  format: format.combine(
    format.timestamp(),
    format.json()
    ),
  defaultMeta: { service: 'disassembler-service' },
  transports: [
    new winston.transports.Console(),
    new winston.transports.File({ filename: 'logs/disassembler.log', timestamp: true })
  ],
});

// Get variable x and equation in the expected form of Ax+Bx+c=0.
// Extracts the coefficients A, b anc C.
export function disassamble(equation, variable) {
    logger.log("info", "Start disassemble equation " + equation);

    const startTime = Date.now();

    var aResult = findA(equation, variable);
    aResult.secondDegreeIndex += 3; // + size of "x^2"

    var bResult = findB(equation, variable, aResult.secondDegreeIndex);
    bResult.firstDegreeIndex += 1; // + size of "x";

    const c = findC(equation, bResult.firstDegreeIndex);

    const endTime = Date.now();

    logger.log("info", "The coefficients of " + equation + " are: " + aResult.a + ", " + bResult.b + ", " + c);
    return new DisassembledEquationMessage(equation, aResult.a, bResult.b, c, startTime, endTime);
}

function findA(equation, variable) {
    const secondDegreeVariable = variable + "^2";
    const secondDegreeIndex = equation.indexOf(secondDegreeVariable, 0);

    if (secondDegreeIndex == -1)
        throw new SyntaxError("Could not find second degree variable");

    logger.log("debug", "Index of " + secondDegreeVariable + ": " + secondDegreeIndex);

    let a = "1";
    if (secondDegreeIndex != 0)
    {
        a = equation.substring(0, secondDegreeIndex);
        if (a[0] == '+')
            a = a.substring(1);
        else if (a[0] != '-' && isNaN(a[0]))
            throw new SyntaxError("Expecting '+' or '-' signs only");

        if (a == "-")
            a = "-1";
    }

    logger.log("debug", "A: " + a);

    return {
        a: a, 
        secondDegreeIndex: secondDegreeIndex,
    }
}

function findB(equation, variable, secondDegreeEndIndex) {
    const firstDegreeIndex = equation.indexOf(variable, secondDegreeEndIndex);

    if (firstDegreeIndex == -1)
        throw new SyntaxError("Could not find first degree variable");

    logger.log("debug", "Index of " + variable + ": " + firstDegreeIndex);

    let b = equation.substring(secondDegreeEndIndex, firstDegreeIndex);
    if (b[0] == '+')
        b = b.substring(1);
    else if (b[0] != '-')
        throw new SyntaxError("Expecting '+' or '-' signs only");

    if (b == "")
        b = "1";

    logger.log("debug", "B: " + b);

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

    logger.log("debug", "C: " + c);
    return c;
}

export function createMessage(equationCoefficients) {
    const jsonMessage = JSON.stringify(equationCoefficients)
    logger.log("info", jsonMessage);
}