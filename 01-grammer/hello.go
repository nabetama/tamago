package main

import (
	"errors"
	"fmt"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			// runtime error: index out of range
			fmt.Println("recover!")
		}
	}()

	a := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		panic(errors.New("index out of range."))
	}
}
