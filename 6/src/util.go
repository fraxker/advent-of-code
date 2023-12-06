package src

import "math"

func Solve_Quadratic(a, b, c float64) []float64 {
	discrimimiant := b*b - 4*a*c
	

	intersection := make([]float64, 0, 2)
	if discrimimiant > 0 {
		sqrt_discriminant := math.Sqrt(discrimimiant)
		x1 := (-b + sqrt_discriminant) / (2 * a)
		x2 := (-b - sqrt_discriminant) / (2 * a)
		intersection = append(intersection, x1, x2)
	} else if discrimimiant == 0 {
		x := -b / (2 * a)
		intersection = append(intersection, x)
	}
	return intersection
}