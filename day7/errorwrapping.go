package day7

import (
	"errors"
	"fmt"
	"os"
)

var ErrFileMissing = errors.New("File is missing")

func Openfile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("%W :%v", ErrFileMissing, err)
		}
		return nil, err
	}
	return file, nil

}

func Errorwrapping() {
	_, err := Openfile("unknown.txt")
	if err != nil {
		if errors.Is(err, ErrFileMissing) {
			fmt.Println("Main detected , the file does not exists")
		} else {
			fmt.Println("Some other errors", err)
		}
	}
}
