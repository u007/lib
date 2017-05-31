package tools

import (
	"fmt"
	"math"
	"strconv"
)

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func RoundAsInt(val float64) int {
	return int(val + math.Copysign(0.5, val))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(RoundAsInt(num*output)) / output
}

func ParseIntFromString(str string, default_value int) (int, error) {
	if str == "" {
		return default_value, nil
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return default_value, err
	}
	return i, nil
}

func ParseIntFromBytes(b []byte) (int, error) {
	s := fmt.Sprintf("%d", b)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func ParseIntFromByte(b byte) (int, error) {
	s := fmt.Sprintf("%d", b)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
