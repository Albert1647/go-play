package utils

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, str := range strings {
		float, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, errors.New("Fail")
		}
		floats = append(floats, float)
	}
	return floats, nil
}
