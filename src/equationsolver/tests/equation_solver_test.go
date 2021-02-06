package equationsolvertests

import (
	"equationsolver"
	"testing"
	"strconv"
)

func TestFindRoots_EquationWithRoots_RootsAreCorrect(t *testing.T) {

	testSources := [][]string {
		{"1", "-5", "6", "3", "2"},
		{"2", "4", "-6", "1", "-3"},
	}

	for _, testSource := range testSources {
		a, _ := strconv.Atoi(testSource[0])
		b, _ := strconv.Atoi(testSource[1])
		c, _ := strconv.Atoi(testSource[2])
		expectedRoot1, _ := strconv.ParseFloat(testSource[3], 32)
		expectedRoot2, _ := strconv.ParseFloat(testSource[4], 32)
		
		root1, root2, hasRoots := equationsolver.FindRoots(a, b, c)

		if !hasRoots {
			t.FailNow()
		}

		if expectedRoot1  != root1 {
			t.FailNow()
		}

		if expectedRoot2  != root2 {
			t.FailNow()
		}
	}
}

func TestFindRoots_EquationWithOneRoot_RootIsCorrect(t *testing.T) {

	testSources := [][]string {
		{"-4", "12", "-9", "1.5"},
	}

	for _, testSource := range testSources {
		a, _ := strconv.Atoi(testSource[0])
		b, _ := strconv.Atoi(testSource[1])
		c, _ := strconv.Atoi(testSource[2])
		expectedRoot1, _ := strconv.ParseFloat(testSource[3], 32)
		
		root1, root2, hasRoots := equationsolver.FindRoots(a, b, c)

		if !hasRoots {
			t.FailNow()
		}

		if expectedRoot1  != root1 {
			t.FailNow()
		}

		if root1 != root2 {
			t.FailNow()
		}
	}
}

func TestFindRoots_EquationWithoutRoots_ReturnsNoResult(t *testing.T) {

	testSources := [][]string {
		{"1", "-3", "4"},
	}

	for _, testSource := range testSources {
		a, _ := strconv.Atoi(testSource[0])
		b, _ := strconv.Atoi(testSource[1])
		c, _ := strconv.Atoi(testSource[2])
		
		_, _, hasRoots := equationsolver.FindRoots(a, b, c)

		if hasRoots {
			t.FailNow()
		}
	}
}

func TestFindRoots_EquationWithZeroACoefficientOneRoot_ReturnsNoResult(t *testing.T) {

	testSources := [][]string {
		{"0", "-3", "4"},
	}

	for _, testSource := range testSources {
		a, _ := strconv.Atoi(testSource[0])
		b, _ := strconv.Atoi(testSource[1])
		c, _ := strconv.Atoi(testSource[2])
		
		_, _, hasRoots := equationsolver.FindRoots(a, b, c)

		if hasRoots {
			t.FailNow()
		}
	}
}