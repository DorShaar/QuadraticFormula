package equationarranger

import (
	"equationscanner"
	"log"
	"strings"
	"strconv"
)
	
type scanner interface {
    Scan() error
}

type EquationArranger struct {
	EquationScanner 	equationscanner.EquationScanner
}

func (equationArranger *EquationArranger)Arrange(equation string) ArrangeResult {
    log.Printf("Arranging equation %s", equation);

	err := equationArranger.EquationScanner.Scan(equation)

	if err != nil {
		log.Printf("Arranging equation %s failed with error %s", equation, err);
        return &FailedArrangeResult { originalEquation: equation }
	}

    return innerArrange(equationArranger, equation);
}

func innerArrange(equationArranger *EquationArranger, equation string) ArrangeResult {
    a := getCoefficientOfXSquare(equationArranger)
    b := getCoefficientOfX(equationArranger)
    c := getFreeNumber(equationArranger)

    log.Printf("Coefficients are a: %d b: %d c:%d", a, b, c);

    equationBuilder := strings.Builder{}

    if a != 1 {
        equationBuilder.WriteString(strconv.Itoa(a))
    }

    equationBuilder.WriteString("x^2")

    if b >= 0 {
        equationBuilder.WriteString("+")
    }

    if b != 1 {
        equationBuilder.WriteString(strconv.Itoa(b))
    }

    equationBuilder.WriteString("x")

    if c >= 0 {
        equationBuilder.WriteString("+")
    }

    equationBuilder.WriteString(strconv.Itoa(c))
    equationBuilder.WriteString("=0")

	arrangedEquation := equationBuilder.String()

    log.Printf("Arranged equation of %s is: %s", equation, arrangedEquation)

 	return &SuccessArrangeResult { originalEquation: equation, arrangedEquation: arrangedEquation }
}

func getCoefficientOfXSquare(equationArranger *EquationArranger) int {
    a := 0

    for i := 0; i < len(equationArranger.EquationScanner.SecondDegreeCoefficientsOnLeft); i++ {
        a += equationArranger.EquationScanner.SecondDegreeCoefficientsOnLeft[i]
    }

    for i := 0; i < len(equationArranger.EquationScanner.SecondDegreeCoefficientsOnRight); i++ {
        a -= equationArranger.EquationScanner.SecondDegreeCoefficientsOnRight[i]
    }

    return a
}

func getCoefficientOfX(equationArranger *EquationArranger) int {
    a := 0

    for i := 0; i < len(equationArranger.EquationScanner.FirstDegreeCoefficientsOnLeft); i++ {
        a += equationArranger.EquationScanner.FirstDegreeCoefficientsOnLeft[i]
    }

    for i := 0; i < len(equationArranger.EquationScanner.FirstDegreeCoefficientsOnRight); i++ {
        a -= equationArranger.EquationScanner.FirstDegreeCoefficientsOnRight[i]
    }

    return a
}

func getFreeNumber(equationArranger *EquationArranger) int {
    a := 0

    for i := 0; i < len(equationArranger.EquationScanner.FreeNumbersOnLeft); i++ {
        a += equationArranger.EquationScanner.FreeNumbersOnLeft[i]
    }

    for i := 0; i < len(equationArranger.EquationScanner.FreeNumbersOnRight); i++ {
        a -= equationArranger.EquationScanner.FreeNumbersOnRight[i]
    }

    return a
}