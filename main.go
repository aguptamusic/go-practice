package main

import (
	"fmt"
	//syntax "go-practice/syntax"
	//greetings "go-practice/greetings"
	algos "go-practice/simple-algos"
)

func main() {
	// syntax.Syntax()
	// greetings.SayHello()

	//binary search
	a := [5]int{1, 2, 3, 4, 5}
	index := algos.BinarySearch(a, 3)
	fmt.Println("Found at index:", index)

	//valid parens
	valid := algos.ValidParens("{[]}")
	fmt.Println("Valid:", valid)
	fmt.Println("Expect: true")
	valid = algos.ValidParens("([)]")
	fmt.Println("Valid:", valid)
	fmt.Println("Expect: false")
}
