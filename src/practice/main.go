package main

import (
	"fmt"

	"github.com/priyankitshukla/go/src/practicepackage"
)

func main() {

	fmt.Println("Hello world")
	practicepackage.Hello()
	sum := practicepackage.Sum(5, 5)
	fmt.Println(sum)
}
