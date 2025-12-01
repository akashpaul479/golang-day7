package day7

import (
	"errors"
	"fmt"
	"math"
)

func squareroot(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("cannot calculate squareroot of a negitive number")

	}
	return math.Sqrt(n), nil
}
func Errorhandling() {
	value, err := squareroot(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", value)
	}

	data, err := squareroot(-9)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", data)
	}
}
