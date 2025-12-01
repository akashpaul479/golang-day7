package day7

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidAge   = errors.New("Invalid age")
	ErrInvalidMarks = errors.New("Invalid marks")
)

type students struct {
	Name  string
	Age   int
	marks int
}

func Validstudents(s students) error {
	if s.Age < 0 || s.Age > 100 {
		return fmt.Errorf("Student age error :%w", ErrInvalidAge)
	}
	if s.marks < 0 || s.marks > 100 {
		return fmt.Errorf("Student marks error :%w", ErrInvalidMarks)
	}
	return nil
}
func Customerror1() {
	s := students{Name: "Akash", Age: -5, marks: 120}
	err := Validstudents(s)
	if err != nil {
		if errors.Is(err, ErrInvalidAge) {
			fmt.Println("Error :Age must be between 0 to 100")
		}
		if errors.Is(err, ErrInvalidMarks) {
			fmt.Println("Error : marks must be between 0 to 100")
		}
		fmt.Println("Full error:", err)
	}
}
