package monkey_test

import (
	"fmt"
	"testing"

	"github.com/sirkon/monkey"
)

func printer() {
	fmt.Println("I am one")
}

func TestSimple(t *testing.T) {
	printer()
	monkey.Patch(printer, func() {
		fmt.Println("not what you think")
	})
	printer()
}
