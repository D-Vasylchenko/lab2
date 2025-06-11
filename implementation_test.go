package lab2

import (
	"fmt"
	"testing"

	_ "gopkg.in/check.v1"
)

func TestPrefixToPostfix(t *testing.T) {
	
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
