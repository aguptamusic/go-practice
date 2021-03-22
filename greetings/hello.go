package greetings

import (
  "fmt"
)

func SayHello() {
	// no need for import in same package
  message := Hello("Avi")
  fmt.Println(message)
}