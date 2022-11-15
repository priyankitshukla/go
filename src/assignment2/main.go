package main

import "fmt"

func main() {
	var year int
	var number int
	fmt.Println("Please enter your birth year in YYYY format : ")
	fmt.Scanln(&year)
	isLeapYear(year)
	fmt.Println("Please enter a number : ")
	fmt.Scanln(&number)
	checkEvenOrOdd(number)
	var num1, num2 int
	fmt.Println("Please enter first  number : ")
	fmt.Scanln(&num1)
	fmt.Println("Please enter second  number : ")
	fmt.Scanln(&num2)
	sum(num1, num2)
	minus(num1, num2)
	division(num1, num2)
	mul(num1, num2)
	fmt.Println("Enter a string ")

	var text string
	fmt.Scan(&text)
	textArray := []string{text}
	fmt.Println("Length of entered string is ", len(textArray))
	fmt.Println("Last char of entered string is ", textArray[len(textArray)-1])
	fmt.Println("Enter a number to find ASCII value")
	var ascii int
	fmt.Scan(&ascii)
	fmt.Printf("Enter number is %d and ascii value is %s", ascii, string(ascii))
	reverseString(text)
	var str1 string
	var str2 string
	fmt.Println("Enter two string seperated by comma")
	fmt.Scan(&str1, &str2)
	concatString(str1, str2)

}

func concatString(str1 string, str2 string) {
	fmt.Println(str1 + str2)
}

func reverseString(text string) {
	R := make([]string, len(text))
	count := 0
	for i := len(text) - 1; i >= 0; i-- {
		R[count] = string(text[i])
		count++
	}
	fmt.Printf("Reverse of string %s is %s", text, R)
}

func mul(num1 int, num2 int) {
	fmt.Println("multiplication of two number is \n", num1*num2)
}

func division(num1 int, num2 int) {
	fmt.Println("division of two number is \n", num1/num2)
}

func minus(num1 int, num2 int) {
	fmt.Println("minus of two number is \n", num1-num2)
}

func sum(num1 int, num2 int) {
	fmt.Println("sum of two number is \n", num1+num2)
}
func checkEvenOrOdd(number int) {
	if number%2 == 0 {
		fmt.Printf("Entered number %d is Even number\n", number)
	} else {
		fmt.Printf("Entered number %d is Odd number\n", number)

	}

	if number > 0 {
		fmt.Printf("Entered number %d is Positive number\n", number)
	} else {
		fmt.Printf("Entered number %d is Negative number\n", number)

	}

	if number%10 == 0 {
		fmt.Printf("Entered number %d is single digit number\n", number)
	} else {
		fmt.Printf("Entered number %d is multi digit number\n", number)

	}

}
func isLeapYear(year int) {
	if year%4 == 0 {
		fmt.Println("Year ", year, " is a Leap year")
	} else {
		fmt.Println("Year ", year, " is not a Leap year")

	}
}
