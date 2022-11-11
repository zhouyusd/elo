package utils

import "math"

func CalcExpectation(ra, rb float64) float64 {
	return 1.0 / (1.0 + math.Pow(10.0, (rb-ra)/400.0))
}

func CalcDelta(e, s, k float64) float64 {
	return k * (s - e)
}
