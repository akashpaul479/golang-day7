package day7

import "fmt"

type notfounderror struct {
	Name string
}

func (e notfounderror) Error() string {
	return fmt.Sprintf("%s not found", e.Name)
}

func findstudents(name string, list []string) error {
	for _, s := range list {
		if s == name {
			return nil
		}
	}
	return &notfounderror{Name: name}
}

func Customerror() {
	err := findstudents("Akash", []string{"Bunty", "kushal"})
	if err != nil {
		if nf, ok := err.(*notfounderror); ok {
			fmt.Println("custom error", nf)
		}
	}
}
