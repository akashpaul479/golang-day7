package project

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidAge   = errors.New("Invalid age")
	ErrInvalidMarks = errors.New("Invalid marks")
	ErrDuplicate    = errors.New("Duplicate student")
	ErrNotFound     = errors.New("Students not found")
)

type student struct {
	Name  string
	Age   int
	Marks int
}
type ValidationError struct {
	problems []error
}

func (v ValidationError) Error() string {
	if len(v.problems) == 0 {
		return "no validation error"
	}
	parts := make([]string, 0, len(v.problems))
	for _, p := range v.problems {
		parts = append(parts, p.Error())
	}
	return "Validation error:" + strings.Join(parts, "; ")
}
func (v ValidationError) Is(target error) bool {
	for _, p := range v.problems {
		if errors.Is(p, target) {
			return true
		}
		if ve, ok := p.(ValidationError); ok {
			if ve.Is(target) {
				return true
			}
		}
	}
	return false
}
func (v *ValidationError) add(err error) {
	v.problems = append(v.problems, err)
}
func validatestudent(s student) error {
	var ve ValidationError
	if strings.TrimSpace(s.Name) == "" {
		ve.add(fmt.Errorf("name empty"))
	}
	if s.Age < 0 || s.Age > 100 {
		ve.add(fmt.Errorf("age %d id out of range %w", s.Age, ErrInvalidAge))
	}
	if s.Marks < 0 || s.Marks > 100 {
		ve.add(fmt.Errorf("marks %d is out of range %w", s.Marks, ErrInvalidMarks))
	}
	if len(ve.problems) > 0 {
		return ve
	}
	return nil
}

type studentdatabase struct {
	store map[string]student
}

func newstudentdatabase() *studentdatabase {
	return &studentdatabase{store: make(map[string]student)}
}
func (db *studentdatabase) Addstudents(s student) error {
	if err := validatestudent(s); err != nil {
		return fmt.Errorf("Failed to add students %q :%w", s.Name, err)
	}
	if _, ok := db.store[s.Name]; ok {
		return fmt.Errorf("failed to add students %q :%w ", s.Name, ErrDuplicate)
	}
	db.store[s.Name] = s
	return nil
}
func (db *studentdatabase) Getstudent(name string) (student, error) {
	s, ok := db.store[name]
	if !ok {
		return student{}, fmt.Errorf("get student %q :%w", name, ErrNotFound)
	}
	return s, nil
}
func Reporterror(err error) {
	if err == nil {
		fmt.Println("no error")
		return
	}
	fmt.Println("Error:", err)
	if errors.Is(err, ErrDuplicate) {
		fmt.Println("Reason duplicate students(already exists)")
	}
	if errors.Is(err, ErrInvalidAge) {
		fmt.Println("Readon invalid age (must be between 0 to 100)")
	}
	if errors.Is(err, ErrInvalidMarks) {
		fmt.Println("Readon invalid marks (must be between 0 to 100)")
	}
	var ve ValidationError
	if errors.As(err, &ve) {
		fmt.Println("__validation error__")
		for i, p := range ve.problems {
			fmt.Printf("  %d) %s\n", i+1, p.Error())
		}
	}
}
func StudentValidator() {
	db := newstudentdatabase()
	cases := []student{
		{"Akash", 20, 88},
		{"", 20, 89},
		{"Bunty", -5, 67},
		{"kushal", 30, 150},
		{"Akash", 24, 67},
	}
	for _, s := range cases {
		fmt.Printf("Adding students: %+v\n", s)
		if err := db.Addstudents(s); err != nil {
			Reporterror(err)
		} else {
			fmt.Println("Added succesfully")
		}
	}
	fmt.Println("__Fetch examples__")
	if st, err := db.Getstudent("Akash"); err != nil {
		Reporterror(err)
	} else {
		fmt.Printf("Found student Akash: %+v\n", st)
	}
	if _, err := db.Getstudent("No such"); err != nil {
		Reporterror(err)
	}
}
