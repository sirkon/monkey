package main

import (
	"fmt"

	"github.com/sirkon/monkey"
)

func zero() int {
	return 0
}

func main() {
	monkey.Patch(zero, func() int {
		return 1
	})

	v := zero()
	fmt.Println(v)
}
