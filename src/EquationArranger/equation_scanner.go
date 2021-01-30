package equationscanner

import (
	"log"
	"unicode"
	"strconv"
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
func (equationScanner *EquationScanner) Scan(equation string) {
	initializeScanner(equationScanner)
    collectCoefficients(equationScanner, equation);
    end(equationScanner);
}

func initializeScanner(equationScanner *EquationScanner) {
    equationScanner.SecondDegreeCoefficientsOnLeft = make([]int, 0)
    equationScanner.FirstDegreeCoefficientsOnLeft = make([]int, 0)
    equationScanner.FreeNumbersOnLeft = make([]int, 0)
    equationScanner.SecondDegreeCoefficientsOnRight = make([]int, 0)
    equationScanner.FirstDegreeCoefficientsOnRight = make([]int, 0)
    equationScanner.FreeNumbersOnRight = make([]int, 0)
    equationScanner.VariableSign = ' ';

    equationScanner.isAfterEqualSign = false;
    equationScanner.scanState = DuringStart;
    equationScanner.charactersGroup = make([]rune, 0)
    equationScanner.coefficientSign = '+';
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
			handleDigit(equationScanner, char);
            continue;
		}

		if (unicode.IsLetter(char)) {
            handleVariable(char);
            continue;
        }

        handleSignCharacter(char);
    }

    handleScanEnd();
}

func validateCharacterWithState(equationScanner *EquationScanner, char rune) {
    switch equationScanner.scanState {
        case DuringStart:
        case DuringEqualSign:
            if unicode.IsDigit(char) || unicode.IsLetter(char) || char == '-' || char == '+' {
                return
            }

            break

        case DuringNumber:
            if unicode.IsDigit(char) || unicode.IsLetter(char) || char == '-' || 
            char == '+' || char == '*' || char == '=' {
           		return
            }
                

            break

        case DuringCoefficientSign:
            if unicode.IsDigit(char) || unicode.IsLetter(char) {
                return
            }

            break

        case DuringMultiplySign:
            if unicode.IsLetter(char) {
                return
            }

            break

        case DuringVariable:
            if char == '-' || char == '+' || char == '=' || char == '^' {
                return
            }

            break

        case DuringExponentSign:
            if unicode.IsDigit(char) {
                return
            }

            break

        case DuringExponentNumber:
            if char == '-' || char == '+' {
                return
            }

            break

        log.Panicf("Unexpected case encountered %s", equationScanner.scanState)
    }

    log.Panicf("Invalid position for character '%s. Current State: %s", char, equationScanner.scanState)
}

func handleDigit(equationScanner *EquationScanner, char rune) {
	if (equationScanner.scanState == DuringExponentSign) {
        handleEndOfExponentScan(equationScanner, char);
        equationScanner.scanState = DuringExponentNumber;
        return;
    }

    equationScanner.scanState = DuringNumber;
    equationScanner.charactersGroup = append(equationScanner.charactersGroup, char);
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

func addCoefficientOfExponent2(equationScanner *EquationScanner) {
    number := CreateNumberFromCharactersGroupAndCoefficientSign(equationScanner)
    if equationScanner.isAfterEqualSign {
        equationScanner.SecondDegreeCoefficientsOnRight = append(equationScanner.SecondDegreeCoefficientsOnRight, number)
        return;
    }

    equationScanner.SecondDegreeCoefficientsOnLeft = append(equationScanner.SecondDegreeCoefficientsOnLeft, number)
}

func addCoefficient(equationScanner *EquationScanner) {
    number := CreateNumberFromCharactersGroupAndCoefficientSign(equationScanner)
    if equationScanner.isAfterEqualSign {
    	equationScanner.FirstDegreeCoefficientsOnRight = append(equationScanner.FirstDegreeCoefficientsOnRight, number)
        return;
    }

	equationScanner.FirstDegreeCoefficientsOnLeft = append(equationScanner.FirstDegreeCoefficientsOnLeft, number)
}

func addFreeNumber(equationScanner *EquationScanner) {
    number := CreateNumberFromCharactersGroupAndCoefficientSign(equationScanner)
    if equationScanner.isAfterEqualSign {
        equationScanner.FreeNumbersOnRight = append(equationScanner.FreeNumbersOnRight, number)
        return;
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
    number := 0;
    for i := 0; i < len(equationScanner.charactersGroup); i++ {
		number = number * 10;
		numericValue, _ := strconv.Atoi(equationScanner.charactersGroup[i])
        number += numericValue;
	}

    equationScanner.charactersGroup = nil
    equationScanner.charactersGroup = make([]rune, 0)

    return number
}

    // private void handleVariable(char ch) throws EquationScanException
    // {
    //     if (scanState == ScanState.DuringCoefficientSign ||
    //         scanState == ScanState.DuringEqualSign ||
    //         scanState == ScanState.DuringStart)
    //     {
    //         charactersGroup.add('1');
    //     }

    //     scanState = ScanState.DuringVariable;

    //     if (variableSign == ' ')
    //     {
    //         variableSign = ch;
    //         return;
    //     }

    //     if (variableSign == ch)
    //         return;

    //     throw new EquationScanException("Found two different variables in the equation. " +
    //             "Variables: " + variableSign + ", " + ch);
    // }

    // private void handleSignCharacter(char ch) throws EquationScanException
    // {
    //     switch (ch)
    //     {
    //         case '^':
    //             handleExponentSign();
    //             break;

    //         case '-':
    //         case '+':
    //             handleCoefficientSign(ch);
    //             break;

    //         case '=':
    //             handleEqualSign(ch);
    //             break;

    //         case '*':
    //             handleMultiplySign(ch);
    //             break;

    //         default:
    //             throw new EquationScanException("Unexpected character '" + ch + "'");
    //     }
    // }

    // private void handleExponentSign()
    // {
    //     scanState = ScanState.DuringExponentSign;
    // }

    // private void handleCoefficientSign(char ch) throws EquationScanException
    // {
    //     if (scanState == ScanState.DuringNumber)
    //         addFreeNumber();

    //     if (scanState == ScanState.DuringVariable)
    //         addCoefficient();

    //     scanState = ScanState.DuringCoefficientSign;
    //     coefficientSign = ch;
    // }

    // private void handleEqualSign(char ch) throws EquationScanException
    // {
    //     if (isAfterEqualSign)
    //         throw new EquationScanException("Equation has two '=' signs");

    //     if (scanState == ScanState.DuringVariable)
    //         addCoefficient();

    //     if (scanState == ScanState.DuringNumber)
    //         addFreeNumber();

    //     scanState = ScanState.DuringEqualSign;
    //     isAfterEqualSign = true;
    // }

    // private void handleMultiplySign(char ch)
    // {
    //     scanState = ScanState.DuringMultiplySign;
    // }

    // private void handleScanEnd() throws EquationScanException
    // {
    //     switch (scanState)
    //     {
    //         case DuringExponentNumber:
    //             return;

    //         case DuringNumber:
    //             addFreeNumber();
    //             break;

    //         case DuringVariable:
    //             addCoefficient();
    //             break;

    //         default:
    //             throw new EquationScanException("Invalid end of equation");
    //     }
    // }