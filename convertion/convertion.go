package convertion

import (
	"errors"
	"strconv"
)

func StringsToFloats(listStrings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range listStrings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("Converting string to float failed: " + err.Error())
		}
		floats = append(floats, floatVal)
	}
	return floats, nil
}