package equationdisassemblertests

import (
	"equationdisassembler"
	"testing"
)

func TestDisassamble_Valid_Equations_As_Expected(t *testing.T) {
	testSources := [][]string {
		{"-36y^2+385y+3=0", "y", "-36", "385", "3"},
		{"3x^2+x+3=0", "x", "3", "1", "3"},
		{"x^2+2x+3=0", "x", "1", "2", "3"},
		{"-x^2+2x+3=0", "x", "-1", "2", "3"},
		{"-36z^2+385z+3=0", "z", "-36", "385", "3"},
		{"36z^2-385z-3=0", "z", "36", "-385", "-3"},
	}

	for _, testSource := range testSources {
		equation := testSource[0]
		variable := testSource[1]
		disassembleMessage := equationdisassembler.Disassemble(equation, variable)

		if  testSource[2] != disassembleMessage.A {
			t.FailNow()
		}

		if  testSource[3] != disassembleMessage.B {
			t.FailNow()
		}

		if  testSource[4] != disassembleMessage.C {
			t.FailNow()
		}

		if  equation != disassembleMessage.Equation {
			t.FailNow()
		}
	}
}

func TestSDisassamble_Invalid_Equations_Returns_Failure(t *testing.T) {
	testSources := [][]string {
		{"-36x^2+385y+3=0", "x"},
		{"-36y^2+385y+3=0", "x"},
		{"-36y^2+385y+3", "y"},
		{"-36y^2 +385y+3=0", "y"},
	}

	for _, testSource := range testSources {
		equation := testSource[0]
		variable := testSource[1]
		disassembleMessage := equationdisassembler.Disassemble(equation, variable)

		if  !disassembleMessage.IsDisassembleFailed {
			t.FailNow()
		}
	}
}