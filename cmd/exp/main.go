package main

import (
	"errors"
	"fmt"
)

func main() {
	err := B()
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("error was of type %v\n", err)
	}
}

var ErrNotFound = errors.New("not found")

func A() error {
	return ErrNotFound
}

func B() error {
	err := A()
	if err != nil {
		return fmt.Errorf("b: %w", err)
	}
	return nil
}
