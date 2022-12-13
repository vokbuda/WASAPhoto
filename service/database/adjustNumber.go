package database

import (
	"math"
	"strconv"
)

func adjustNumber(num_entities uint64) string {
	if num_entities >= 1000000 {

		return strconv.FormatFloat(toFixed(float64(num_entities)/1000000.0), 'E', -1, 64) + "m"
	}
	if num_entities >= 1000 {
		return strconv.FormatFloat(toFixed(float64(num_entities)/1000000.0), 'E', -1, 64) + "k"
	}
	return strconv.Itoa(int(num_entities))
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
func toFixed(num float64) float64 {
	output := math.Pow(10, float64(2))
	return float64(round(num * output))
}
