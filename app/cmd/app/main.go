package main

import (
	"fmt"
)

func ok() (a int) {
	a = 1

	defer func() {
		a = 2
	}()

	return a
}

func main() {
	print(ok())

	fmt.Println([]byte(string(1212)))
}
