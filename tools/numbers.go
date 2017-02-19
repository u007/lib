package tools

import (
	"fmt"
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
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
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
