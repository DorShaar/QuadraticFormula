package equationdisassembler

import (
	"fmt"
	"log"
	"strings"
	"errors"
	"unicode"
)

const secondDegreeSize = 3 // size of "x^2"

// ReadCsv reads from csv path.
func Disassamble(equation string, variable string) DisassembledEquationMessage {
    log.Printf("Start disassemble equation %s", equation)

    a, secondDegreeIndex, err := findA(equation, variable)

    if err != nil {
		log.Print(err)
    }

    secondDegreeIndex += secondDegreeSize 

    // var bResult = findB(equation, variable, aResult.secondDegreeIndex);
    // bResult.firstDegreeIndex += 1; // + size of "x";

    // const c = findC(equation, bResult.firstDegreeIndex);

    // const endTime = Date.now();

    // logger.log("info", "The coefficients of " + equation + " are: " + aResult.a + ", " + bResult.b + ", " + c);
    // return new DisassembledEquationMessage(equation, aResult.a, bResult.b, c, startTime, endTime);
}

func findA(equation string, variable string) (int, int, error) {
    secondDegreeVariable := variable + "^2"
    secondDegreeIndex := strings.Index(equation, secondDegreeVariable)

    if secondDegreeIndex <= -1 {
    	errorMessage := fmt.Sprintf("Could not find second degree variable for equation %s", equation)
        return -1, -1, errors.New(errorMessage)
    }

    log.Printf("Index of %s: %d", secondDegreeVariable, secondDegreeIndex)

    a := "1";
    if secondDegreeIndex != 0 {
		runes := []rune(equation)
	    runes = runes[0:secondDegreeIndex])

        if runes[0] == '+' {
            a = runes[1:]
        }
        else if a[0] != '-' && !unicode.IsDigit(a[0]) {
        	return -1, -1, errors.New("Expecting '+' or '-' signs only")
        }

        if (a == "-")
            a = "-1"
    }

    log.Printf("Found A: %d", a)

    return a, secondDegreeIndex, nil
}