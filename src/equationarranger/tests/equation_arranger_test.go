package equationarrangertests

import (
	"equationarranger"
	"equationscanner"
	"testing"
)

func TestArrange_ValidEquations_ArrangedAsExpected(t *testing.T) {
	testSources := [][]string {
		{"4x^2 - 22 =x^2+5", "3x^2+0x-27=0"},
        {"4x^2- 22", "4x^2+0x-22=0"},
        {"-5x + 3", "0x^2-5x+3=0"},
        {"y ^ 2 + y  = 1", "x^2+x-1=0"},
	}

	equationScanner := equationscanner.EquationScanner{}
	equationArranger := equationarranger.EquationArranger{ EquationScanner: equationScanner}

	for _, testSource := range testSources {
		equation := testSource[0]
		expectedArrangedEquation := testSource[1]

		arrangeResult := equationArranger.Arrange(equation)

        if !arrangeResult.IsArrangeSucceeded() {
        	t.FailNow()
        }

        if arrangeResult.ArrangedEquation() != expectedArrangedEquation {
        	t.FailNow()
        }
	}
}

func TestArrange_InvalidEquations_ArrangedNotSucceed(t *testing.T) {
	invalidEquations := []string {
		"abc",
		"y ^ 2 + x  = 1",
		"x^3 + x^2 + 1 = 1",
	}

	equationScanner := equationscanner.EquationScanner{}
	equationArranger := equationarranger.EquationArranger{ EquationScanner: equationScanner}

	for _, invalidEquation := range invalidEquations {
		arrangeResult := equationArranger.Arrange(invalidEquation)

        if arrangeResult.IsArrangeSucceeded() {
        	t.FailNow()
        }
	}
}