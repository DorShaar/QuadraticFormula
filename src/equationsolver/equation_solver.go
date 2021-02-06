package equationsolver

import (
	"log"
	"math"
)

// FindRoots get 3 coefficients of quadratic equation, a,b,c, and return the two root and bool if they were found.
func FindRoots(a int, b int, c int) (float64, float64, bool) {
	log.Printf("Calcualting roots for coefficient a: %d, b: %d, c: %d", a, b, c)

	if a == 0 {
		log.Printf("Coefficients a: %d, b: %d, c: %d has no quadratic root since a = 0", a, b, c)
        return 0, 0, false
    }

    b_float64 := float64(b)

    delta := math.Pow(b_float64, 2) - float64(4 * a * c);

    if delta < 0 {
    	log.Printf("Coefficients a: %d, b: %d, c: %d has no quadratic root since has no quadratic roots. Delta = %f", a, b, c, delta)
        return 0, 0, false
    }

    denominator_float := float64(2 * a)

    x1 := (-b_float64 + math.Sqrt(delta)) / denominator_float;
    x2 := (-b_float64 - math.Sqrt(delta)) / denominator_float;

    return x1, x2, true
}