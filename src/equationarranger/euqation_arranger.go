package equationarranger

import (
	"equationscanner"
	"log"
)
	
type scanner interface {
    Scan() error
}

type EquationArranger struct {
	equationScanner 	equationscanner.EquationScanner
}

func (equationArranger *EquationArranger)Arrange(equation string) ArrangeResult {
    log.Printf("Arranging equation %s", equation);

	err := equationArranger.equationScanner.Scan(equation)

	if err != nil {
		log.Printf("Arranging equation %s failed with error %s", equation, err);
        return &FailedArrangeResult { originalEquation: equation }
	}

	return &FailedArrangeResult { originalEquation: equation }
    // return innerArrange(equation);
}

    // private ArrangeResult innerArrange(String equation)
    // {
    //     int a = getCoefficientOfXSquare();
    //     int b = getCoefficientOfX();
    //     int c = getFreeNumber();

    //     logger.debug("Coefficients are a: " + a + " b: " + b + " c: " + c);

    //     StringBuilder equationBuilder = new StringBuilder();

    //     if (a != 1)
    //         equationBuilder.append(a);

    //     equationBuilder.append("x^2");

    //     if (b >= 0)
    //         equationBuilder.append("+");

    //     if (b != 1)
    //         equationBuilder.append(b);

    //     equationBuilder.append("x");

    //     if (c >= 0)
    //         equationBuilder.append("+");

    //     equationBuilder.append(c).append("=0");

    //     logger.info("Arranged equation of " + equation + " is: " + equationBuilder.toString());

    //     return ArrangeResult.SuccessArrange(equation, equationBuilder.toString());
    // }

    // private int getCoefficientOfXSquare()
    // {
    //     int a = 0;

    //     for (int i = 0; i < equationScanner.secondDegreeCoefficientsOnLeft.size(); ++i)
    //     {
    //         a += equationScanner.secondDegreeCoefficientsOnLeft.get(i);
    //     }

    //     for (int i = 0; i < equationScanner.secondDegreeCoefficientsOnRight.size(); ++i)
    //     {
    //         a -= equationScanner.secondDegreeCoefficientsOnRight.get(i);
    //     }

    //     return a;
    // }

    // private int getCoefficientOfX()
    // {
    //     int a = 0;

    //     for (int i = 0; i < equationScanner.firstDegreeCoefficientsOnLeft.size(); ++i)
    //     {
    //         a += equationScanner.firstDegreeCoefficientsOnLeft.get(i);
    //     }

    //     for (int i = 0; i < equationScanner.firstDegreeCoefficientsOnRight.size(); ++i)
    //     {
    //         a -= equationScanner.firstDegreeCoefficientsOnRight.get(i);
    //     }

    //     return a;
    // }

    // private int getFreeNumber()
    // {
    //     int a = 0;

    //     for (int i = 0; i < equationScanner.freeNumbersOnLeft.size(); ++i)
    //     {
    //         a += equationScanner.freeNumbersOnLeft.get(i);
    //     }

    //     for (int i = 0; i < equationScanner.freeNumbersOnRight.size(); ++i)
    //     {
    //         a -= equationScanner.freeNumbersOnRight.get(i);
    //     }

    //     return a;
    // }