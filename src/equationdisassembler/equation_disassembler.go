package equationdisassembler

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

const secondDegreeSize = 3 // size of "x^2"
const firstDegreeSize = 1 // size of "x"


func Disassemble(equation string, variable string) DisassembledEquationMessage {
	log.Printf("Start disassemble equation %s", equation)

	a, secondDegreeIndex, err := findA(equation, variable)

	if err != nil {
		log.Print(err)
		return DisassembledEquationMessage { IsDisassembleFailed: true }
	}

	b, firstDegreeIndex, err := findB(equation, variable, secondDegreeIndex + secondDegreeSize)

	if err != nil {
		log.Print(err)
		return DisassembledEquationMessage { IsDisassembleFailed: true }
	}

    c, err := findC(equation, firstDegreeIndex + firstDegreeSize)

    if err != nil {
		log.Print(err)
		return DisassembledEquationMessage { IsDisassembleFailed: true }
	}

    log.Printf("The coefficients of %s are: %s, %s, %s", equation, a, b, c)
	return DisassembledEquationMessage { 
        Equation: equation,
        A: a,
        B: b,
        C: c,
    };
}

func findA(equation string, variable string) (string, int, error) {
	secondDegreeVariable := variable + "^2"
	secondDegreeIndex := strings.Index(equation, secondDegreeVariable)

	if secondDegreeIndex <= -1 {
		errorMessage := fmt.Sprintf("Could not find second degree variable for equation %s", equation)
		return "", -1, errors.New(errorMessage)
	}

	log.Printf("Index of %s: %d", secondDegreeVariable, secondDegreeIndex)

	a := "1"
	if secondDegreeIndex != 0 {
		a = equation[0:secondDegreeIndex]

		if a[0] == '+' {
			a = a[1:]
		} else if a[0] != '-' {
			aRune, _ := utf8.DecodeRuneInString(a)
			if !unicode.IsDigit(aRune) {
				return "", -1, errors.New("Expecting '+' or '-' signs only")
			}
		}

		if a == "-" {
			a = "-1"
		}
	}

	log.Printf("Found A: %s", a)

	return a, secondDegreeIndex, nil
}

func findB(equation string, variable string, secondDegreeEndIndex int) (string, int, error) {
	equationAfterSecondDegree := equation[secondDegreeEndIndex:]

	firstDegreeIndex := strings.Index(equationAfterSecondDegree, variable)

	if firstDegreeIndex <= -1 {
		errorMessage := fmt.Sprintf("Could not find first degree variable for equation %s", equation)
		return "", -1, errors.New(errorMessage)
	}

	b := equationAfterSecondDegree[:firstDegreeIndex]

	if b[0] == '+' {
        b = b[1:]
	} else if b[0] != '-' {
		errorMessage := fmt.Sprintf("Expecting '+' or '-' signs only")
		return "", -1, errors.New(errorMessage)
	}

    if len(b) == 0 {
        b = "1"
    }

    firstDegreeIndex += secondDegreeEndIndex
	log.Printf("Index of %s: %d", variable, firstDegreeIndex)
    log.Printf("Found B: %s", b)

    return b, firstDegreeIndex, nil
}

func findC(equation string, firstDegreeEndIndex int) (string, error) {
	equalSignIndex := strings.Index(equation, "=")

    if equalSignIndex == -1 {
    	errorMessage := fmt.Sprintf("Could not find equal sign")
		return "", errors.New(errorMessage)
    }

    c := equation[firstDegreeEndIndex:equalSignIndex]

    if c[0] == '+' {
        c = c[1:]
    }

   	log.Printf("Found C: %s", c)
    return c, nil;
}