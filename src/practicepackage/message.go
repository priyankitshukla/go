package practicepackage

import "fmt"

func Hello() {
	fmt.Println("Message from anoter package")
}

func Sum(a int, b int) int {
	sum := a + b
	return sum
}
