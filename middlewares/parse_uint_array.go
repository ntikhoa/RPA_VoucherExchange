package middlewares

import (
	"errors"
	"strconv"
)

func getUintArrayType(requestForm map[string][]string, key string) ([]uint, error) {
	values, ok := requestForm[key]
	if !ok {
		return nil, errors.New("\"" + key + "\"" + " required")
	}
	var valueUint []uint
	for _, value := range values {
		log.println(value)
		valueUint64, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return valueUint, errors.New("\"" + key + "\" " + " invalid type")
		}
		valueUint = append(valueUint, uint(valueUint64))
	}

	return valueUint, nil
}
