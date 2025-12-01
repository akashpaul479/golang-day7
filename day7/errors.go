package day7

import (
	"errors"
	"fmt"
)

func Errors() {
	err := errors.New("Something went wrong")
	fmt.Println(err)

	num := -9
	err1 := fmt.Errorf("number %d is not valid", num)
	fmt.Println(err1)
}
