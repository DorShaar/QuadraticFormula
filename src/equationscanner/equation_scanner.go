package equationscanner

import (
	"log"
	"unicode"
    "errors"
    "fmt"
)

// EquationScanner scans given equation string and hold the equation's parameters.
type EquationScanner struct {
	SecondDegreeCoefficientsOnLeft 	[]int
	FirstDegreeCoefficientsOnLeft 	[]int
	FreeNumbersOnLeft				[]int
	SecondDegreeCoefficientsOnRight	[]int
	FirstDegreeCoefficientsOnRight	[]int
	FreeNumbersOnRight				[]int
	VariableSign					rune

	isAfterEqualSign				bool
	scanState 						ScanState
	charactersGroup 				[]rune
	coefficientSign 				rune
}

// Connect Creates a connection to given address
func (equationScanner *EquationScanner) Scan(equation string) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = errors.New(fmt.Sprintf("Error while scanning equation %s", equation))
        }
    }()

	initializeScanner(equationScanner)
    collectCoefficients(equationScanner, equation)
    end(equationScanner)

    return
}

func initializeScanner(equationScanner *EquationScanner) {
    equationScanner.SecondDegreeCoefficientsOnLeft = make([]int, 0)
    equationScanner.FirstDegreeCoefficientsOnLeft = make([]int, 0)
    equationScanner.FreeNumbersOnLeft = make([]int, 0)
    equationScanner.SecondDegreeCoefficientsOnRight = make([]int, 0)
    equationScanner.FirstDegreeCoefficientsOnRight = make([]int, 0)
    equationScanner.FreeNumbersOnRight = make([]int, 0)
    equationScanner.VariableSign = ' '

    equationScanner.isAfterEqualSign = false
    equationScanner.scanState = DuringStart
    equationScanner.charactersGroup = make([]rune, 0)
    equationScanner.coefficientSign = '+'
}

func end(equationScanner *EquationScanner) {
	if len(equationScanner.SecondDegreeCoefficientsOnLeft) == 0 {
		equationScanner.SecondDegreeCoefficientsOnLeft = append(equationScanner.SecondDegreeCoefficientsOnLeft, 0)
	}

	if len(equationScanner.FirstDegreeCoefficientsOnLeft) == 0 {
		equationScanner.FirstDegreeCoefficientsOnLeft = append(equationScanner.FirstDegreeCoefficientsOnLeft, 0)
	}

	if len(equationScanner.FreeNumbersOnLeft) == 0 {
		equationScanner.FreeNumbersOnLeft = append(equationScanner.FreeNumbersOnLeft, 0)
	}

	if len(equationScanner.SecondDegreeCoefficientsOnRight) == 0 {
		equationScanner.SecondDegreeCoefficientsOnRight = append(equationScanner.SecondDegreeCoefficientsOnRight, 0)
	}

	if len(equationScanner.FirstDegreeCoefficientsOnRight) == 0 {
		equationScanner.FirstDegreeCoefficientsOnRight = append(equationScanner.FirstDegreeCoefficientsOnRight, 0)
	}

	if len(equationScanner.FreeNumbersOnRight) == 0 {
		equationScanner.FreeNumbersOnRight = append(equationScanner.FreeNumbersOnRight, 0)
	}
}

func collectCoefficients(equationScanner *EquationScanner, equation string) {
	for _, char := range equation {

		if unicode.IsSpace(char) {
			continue
		}

		validateCharacterWithState(equationScanner, char)

		if unicode.IsDigit(char) {
			handleDigit(equationScanner, char)
            continue
		}

		if (unicode.IsLetter(char)) {
            handleVariable(equationScanner, char)
            continue
        }

        handleSignCharacter(equationScanner, char)
    }

    handleScanEnd(equationScanner)
}

func validateCharacterWithState(equationScanner *EquationScanner, char rune) {
    switch equationScanner.scanState {
        case DuringStart, DuringEqualSign:
            if unicode.IsDigit(char) || unicode.IsLetter(char) || char == '-' || char == '+' {
                return
            }

        case DuringNumber:
            if unicode.IsDigit(char) || unicode.IsLetter(char) || char == '-' || 
            char == '+' || char == '*' || char == '=' {
           		return
            }

        case DuringCoefficientSign:
            if unicode.IsDigit(char) || unicode.IsLetter(char) {
                return
            }

        case DuringMultiplySign:
            if unicode.IsLetter(char) {
                return
            }

        case DuringVariable:
            if char == '-' || char == '+' || char == '=' || char == '^' {
                return
            }

        case DuringExponentSign:
            if unicode.IsDigit(char) {
                return
            }

        case DuringExponentNumber:
            if char == '-' || char == '+' {
                return
            }

        log.Panicf("Unexpected case encountered %s", equationScanner.scanState)
    }

    log.Panicf("Invalid position for character '%s. Current State: %s", char, equationScanner.scanState)
}

func handleDigit(equationScanner *EquationScanner, char rune) {
	if (equationScanner.scanState == DuringExponentSign) {
        handleEndOfExponentScan(equationScanner, char)
        equationScanner.scanState = DuringExponentNumber
        return
    }

    equationScanner.scanState = DuringNumber
    equationScanner.charactersGroup = append(equationScanner.charactersGroup, char)
}

func handleEndOfExponentScan(equationScanner *EquationScanner, char rune) {
	defer setScanState(equationScanner, DuringExponentNumber)

	if char == '2' {
		addCoefficientOfExponent2(equationScanner)
		return
	}

	if char == '1' {
		addCoefficient(equationScanner)
		return
	}

	if char == '0' {
		addFreeNumber(equationScanner)
		return
	}

	log.Panicf("Found invalid exponent character %s", char)
}

func setScanState(equationScanner *EquationScanner, newState ScanState) {
	equationScanner.scanState = newState
}

func handleVariable(equationScanner *EquationScanner, char rune) {
    if equationScanner.scanState == DuringCoefficientSign ||
        equationScanner.scanState == DuringEqualSign ||
        equationScanner.scanState == DuringStart {
        equationScanner.charactersGroup = append(equationScanner.charactersGroup, '1')
    }

    equationScanner.scanState = DuringVariable

    if equationScanner.VariableSign == ' ' {
        equationScanner.VariableSign = char
        return
    }

    if equationScanner.VariableSign == char {
        return
    }

    log.Panicf("Found two different variables in the equation. Variables: %s, %s", equationScanner.VariableSign, char)
}

func handleSignCharacter(equationScanner *EquationScanner, char rune) {
    switch char  {
        case '^':
            handleExponentSign(equationScanner)
            break

        case '-', '+':
            handleCoefficientSign(equationScanner, char)
            break

        case '=':
            handleEqualSign(equationScanner, char)
            break

        case '*':
            handleMultiplySign(equationScanner, char)
            break


        log.Panicf("Unexpected character '%s':, %s", equationScanner.VariableSign, char)
    }
}

func handleExponentSign(equationScanner *EquationScanner) {
    equationScanner.scanState = DuringExponentSign
}

func handleCoefficientSign(equationScanner *EquationScanner, char rune) {
    if equationScanner.scanState == DuringNumber {
        addFreeNumber(equationScanner)
    }

    if equationScanner.scanState == DuringVariable {
        addCoefficient(equationScanner)
    }

    equationScanner.scanState = DuringCoefficientSign
    equationScanner.coefficientSign = char
}

func handleEqualSign(equationScanner *EquationScanner, char rune) {
    if equationScanner.isAfterEqualSign{
    	log.Panicf("Equation has two '=' signs")
    }

    if equationScanner.scanState == DuringVariable {
        addCoefficient(equationScanner)
    }

    if equationScanner.scanState == DuringNumber {
        addFreeNumber(equationScanner)
    }

    equationScanner.scanState = DuringEqualSign
    equationScanner.isAfterEqualSign = true
}

func handleMultiplySign(equationScanner *EquationScanner, char rune) {
    equationScanner.scanState = DuringMultiplySign
}

func handleScanEnd(equationScanner *EquationScanner) {
    switch equationScanner.scanState {
        case DuringExponentNumber:
            return

        case DuringNumber:
            addFreeNumber(equationScanner)
            break

        case DuringVariable:
            addCoefficient(equationScanner)
            break

        log.Panic("Invalid end of equation")
    }
}

func addCoefficientOfExponent2(equationScanner *EquationScanner) {
	number := CreateNumberFromCharactersGroupAndCoefficientSign(equationScanner)
	if equationScanner.isAfterEqualSign {
	    equationScanner.SecondDegreeCoefficientsOnRight = append(equationScanner.SecondDegreeCoefficientsOnRight, number)
	    return
	}

    equationScanner.SecondDegreeCoefficientsOnLeft = append(equationScanner.SecondDegreeCoefficientsOnLeft, number)
}

func addCoefficient(equationScanner *EquationScanner) {
    number := CreateNumberFromCharactersGroupAndCoefficientSign(equationScanner)
    if equationScanner.isAfterEqualSign {
    	equationScanner.FirstDegreeCoefficientsOnRight = append(equationScanner.FirstDegreeCoefficientsOnRight, number)
        return
    }

	equationScanner.FirstDegreeCoefficientsOnLeft = append(equationScanner.FirstDegreeCoefficientsOnLeft, number)
}

func addFreeNumber(equationScanner *EquationScanner) {
    number := CreateNumberFromCharactersGroupAndCoefficientSign(equationScanner)
    if equationScanner.isAfterEqualSign {
        equationScanner.FreeNumbersOnRight = append(equationScanner.FreeNumbersOnRight, number)
        return
    }

    equationScanner.FreeNumbersOnLeft = append(equationScanner.FreeNumbersOnLeft, number)
}

func CreateNumberFromCharactersGroupAndCoefficientSign(equationScanner *EquationScanner) int {
    number := CreateNumberFromCharactersGroup(equationScanner)

    defer setCoefficientSign(equationScanner, '+')

    if equationScanner.coefficientSign == '-' {
        return -number
    }

    if equationScanner.coefficientSign == '+' {
        return number
    }

    log.Panicf("Unexpected coefficient sign %s", equationScanner.coefficientSign)
    return 99999999999
}

func setCoefficientSign(equationScanner *EquationScanner, newCoefficientSign rune) {
	equationScanner.coefficientSign = newCoefficientSign
}

func CreateNumberFromCharactersGroup(equationScanner *EquationScanner) int {
    number := 0
    for i := 0; i < len(equationScanner.charactersGroup); i++ {
		number = number * 10
        number += int(equationScanner.charactersGroup[i] - '0')
	}

    equationScanner.charactersGroup = nil
    equationScanner.charactersGroup = make([]rune, 0)

    return number
}