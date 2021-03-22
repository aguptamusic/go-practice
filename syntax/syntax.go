package misc

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

/* Capitalize to make exportable. */
func Syntax() {
	fmt.Println("hello world")
	variables()
	function_call()
	iteration()
	array_and_slice()
	maps()
	structs()
	pointers()
}

func variables() {
	var x int = 5
	y := 7 // shorthand
	fmt.Println("y is:", y)

	// conditionals
	fmt.Println("x is:", x)
	if x > 6 {
		fmt.Println("More than 6.")
	} else if x < 2 {
		fmt.Println("Less than two.")
	} else {
		fmt.Println("2 <= x <= 6")
	}

	// function call
	fmt.Println("sum x+y =", sum(x, y))
}

func sum(x int, y int) int {
	return x + y
}

func function_call() {
	// multiple return values from functions
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sqrt of 16 is", result)
	}

	result, err = sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sqrt of 16 is", result)
	}
}

// can return multiple values!
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("sqrt is undefined for negative numbers.")
	}
	return math.Sqrt(x), nil
}

func iteration() {
	// for loop
	fmt.Println("for loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	fmt.Println("while loop")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func array_and_slice() {
	// array (fixed length)
	var a [5]int
	a[2] = 7

	//shorthand
	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	// slices (array with modifiable size)
	c := []int{5, 4, 3, 2, 1}
	c = append(c, 13) // creates new array bts, doesn't modify old

	// iterate over slice with "range"
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
}

func maps() {
	// maps
	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	fmt.Println("map w/square:", vertices)
	fmt.Println("triangle has", vertices["triangle"], "vertices")
	delete(vertices, "square")
	fmt.Println("map w/o square:", vertices)

	// iterate over map with "range"
	vertices["octagon"] = 8
	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}
}

func structs() {
	p := person{name: "Avi", age: 19}
	fmt.Println(p)
	fmt.Println(p.name, "is", p.age)
}

func pointers() {
	i := 8
	fmt.Println(i, "stored at", &i)
	inc(&i)
	fmt.Println("i is now", i)
}

func inc(x *int) {
	*x++
}
