package equationscannertests

import (
	"equationscanner"
	"testing"
	"strconv"
)

func TestScan_ValidEquations_ScannedAsExpected(t *testing.T) {

	testSources := [][]string {
		{"3x^2-0x-27=0", "3", "0", "-27", "0", "0", "0", "x"},
		{"3y^2-0y-27=0", "3" ,"0", "-27", "0" ,"0", "0", "y"},
        {"36x^2-0x+27=0", "36", "0", "27", "0", "0", "0", "x"},
        {"3*x^2-5*x-27=0", "3", "-5", "-27", "0", "0", "0", "x"},
        {"-3*x^2-0*x-27=4x^2 - 3x", "-3", "0", "-27", "4", "-3", "0", "x"},
        {"-3*x^2-50  *  x-27=4x^2 - 3x  + 2", "-3", "-50", "-27", "4", "-3", "2","x"},
	}

	equationScanner := equationscanner.EquationScanner{}

	for _, testSource := range testSources {
		equationScanner.Scan(testSource[0])

		secondDegreeCoefficientsOnLeft, _ := strconv.Atoi(testSource[1])
		if equationScanner.SecondDegreeCoefficientsOnLeft[0] != secondDegreeCoefficientsOnLeft {
			t.FailNow()
		}

		firstDegreeCoefficientsOnLeft, _ := strconv.Atoi(testSource[2])
		if equationScanner.FirstDegreeCoefficientsOnLeft[0] != firstDegreeCoefficientsOnLeft {
			t.FailNow()
		}

		freeNumbersOnLeft, _ := strconv.Atoi(testSource[3])
		if equationScanner.FreeNumbersOnLeft[0] != freeNumbersOnLeft {
			t.FailNow()
		}

		secondDegreeCoefficientsOnRight, _ := strconv.Atoi(testSource[4])
		if equationScanner.SecondDegreeCoefficientsOnRight[0] != secondDegreeCoefficientsOnRight {
			t.FailNow()
		}

		firstDegreeCoefficientsOnRight, _ := strconv.Atoi(testSource[5])
		if equationScanner.FirstDegreeCoefficientsOnRight[0] != firstDegreeCoefficientsOnRight {
			t.FailNow()
		}

		freeNumbersOnRight, _ := strconv.Atoi(testSource[6])
		if equationScanner.FreeNumbersOnRight[0] != freeNumbersOnRight {
			t.FailNow()
		}

		variableSign := rune(testSource[7][0])
		if equationScanner.VariableSign != variableSign {
			t.FailNow()
		}
	}
}

func TestScan_ValidSpecificEquation1_ScannedAsExpected(t *testing.T) {

	testSources := [][]string {
		{"3x^1-0x-27=0", "0", "3", "0", "-27", "0", "0", "0", "x"},
	}

	equationScanner := equationscanner.EquationScanner{}

	for _, testSource := range testSources {
		equationScanner.Scan(testSource[0])

		secondDegreeCoefficientsOnLeft, _ := strconv.Atoi(testSource[1])
		if equationScanner.SecondDegreeCoefficientsOnLeft[0] != secondDegreeCoefficientsOnLeft {
			t.FailNow()
		}

		firstDegreeCoefficientsOnLeft0, _ := strconv.Atoi(testSource[2])
		if equationScanner.FirstDegreeCoefficientsOnLeft[0] != firstDegreeCoefficientsOnLeft0 {
			t.FailNow()
		}

		firstDegreeCoefficientsOnLeft1, _ := strconv.Atoi(testSource[3])
		if equationScanner.FirstDegreeCoefficientsOnLeft[1] != firstDegreeCoefficientsOnLeft1 {
			t.FailNow()
		}

		freeNumbersOnLeft, _ := strconv.Atoi(testSource[4])
		if equationScanner.FreeNumbersOnLeft[0] != freeNumbersOnLeft {
			t.FailNow()
		}

		secondDegreeCoefficientsOnRight, _ := strconv.Atoi(testSource[5])
		if equationScanner.SecondDegreeCoefficientsOnRight[0] != secondDegreeCoefficientsOnRight {
			t.FailNow()
		}

		firstDegreeCoefficientsOnRight, _ := strconv.Atoi(testSource[6])
		if equationScanner.FirstDegreeCoefficientsOnRight[0] != firstDegreeCoefficientsOnRight {
			t.FailNow()
		}

		freeNumbersOnRight, _ := strconv.Atoi(testSource[7])
		if equationScanner.FreeNumbersOnRight[0] != freeNumbersOnRight {
			t.FailNow()
		}

		variableSign := rune(testSource[8][0])
		if equationScanner.VariableSign != variableSign {
			t.FailNow()
		}
	}
}

func TestScan_ValidSpecificEquation2_ScannedAsExpected(t *testing.T) {

	testSources := [][]string {
		{"3x^0-0x-27=0", "0", "0", "3", "-27", "0", "0", "0", "x"},
	}

	equationScanner := equationscanner.EquationScanner{}

	for _, testSource := range testSources {
		equationScanner.Scan(testSource[0])

		secondDegreeCoefficientsOnLeft, _ := strconv.Atoi(testSource[1])
		if equationScanner.SecondDegreeCoefficientsOnLeft[0] != secondDegreeCoefficientsOnLeft {
			t.FailNow()
		}

		firstDegreeCoefficientsOnLeft, _ := strconv.Atoi(testSource[2])
		if equationScanner.FirstDegreeCoefficientsOnLeft[0] != firstDegreeCoefficientsOnLeft {
			t.FailNow()
		}

		freeNumbersOnLeft0, _ := strconv.Atoi(testSource[3])
		if equationScanner.FreeNumbersOnLeft[0] != freeNumbersOnLeft0 {
			t.FailNow()
		}

		freeNumbersOnLeft1, _ := strconv.Atoi(testSource[4])
		if equationScanner.FreeNumbersOnLeft[1] != freeNumbersOnLeft1 {
			t.FailNow()
		}

		secondDegreeCoefficientsOnRight, _ := strconv.Atoi(testSource[5])
		if equationScanner.SecondDegreeCoefficientsOnRight[0] != secondDegreeCoefficientsOnRight {
			t.FailNow()
		}

		firstDegreeCoefficientsOnRight, _ := strconv.Atoi(testSource[6])
		if equationScanner.FirstDegreeCoefficientsOnRight[0] != firstDegreeCoefficientsOnRight {
			t.FailNow()
		}

		freeNumbersOnRight, _ := strconv.Atoi(testSource[7])
		if equationScanner.FreeNumbersOnRight[0] != freeNumbersOnRight {
			t.FailNow()
		}

		variableSign := rune(testSource[8][0])
		if equationScanner.VariableSign != variableSign {
			t.FailNow()
		}
	}
}

func TestScan_InvalidEquations_ThrowsEquationScanException(t *testing.T) {

	invalidEquations := []string {
		"3xx^0-0x-27=0",
        "3xx^0-0x-27==0",
        "3xx^3-0x-27=0",
        "3xx^22-0x-27=0",
        "3x^2--0x-27=0",
        "3x^2-0**x-27=0",
        "3x^2-0*x-*27=0",
        "3x^2-0*x-27=",
        "3x^2-0y-27=0",
        "abc",
    }

	equationScanner := equationscanner.EquationScanner{}

	for _, invalidEquation := range invalidEquations {
		defer func() {
		        if r := recover(); r == nil {
		            t.Errorf("Equation %s did not paniced", invalidEquation)
		        }
		    }()

		equationScanner.Scan(invalidEquation)
	}
}