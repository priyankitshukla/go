package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Assignment 1 starts here...")
	isprime := isPrime(13)
	fmt.Println(isprime)
	fmt.Println(isPalindrome(121))
	fmt.Println(printSqurt(40))
	fmt.Println(power(4))
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 2, 3}
	fmt.Println("Sum of two array is --->", sumArray(arr1[:], arr2[:]))
	fmt.Println("Mul of two array is --->", mulArray(arr1[:], arr2[:]))
	fmt.Println("Plese enter 3 numbers seperated by space")
	X := make([]int, 3)
	fmt.Scanln(&X[0], &X[1], &X[2])
	reverseArray(X)

}

// WAP to check if number is prime or not
func isPrime(number int) string {

	for i := 2; i < number-1; i++ {
		if number%i == 0 {
			return "Given number is Not a prime number"
		}
	}
	return "Given number is Prime number"
}

func isPalindrome(number int) string {
	input_num := number
	var remainder int
	res := 0
	for number > 0 {
		remainder = number % 10
		res = (res * 10) + remainder
		number = number / 10
		fmt.Println(res, " ---", input_num)

	}
	if input_num == res {
		return "Palindrome"
	} else {
		return "Not a Palindrome"
	}
}

// WAP to find squart of a number

func printSqurt(number int) int {
	return int(math.Sqrt(float64(number)))
}

func power(number float64) int {
	return int(math.Pow(number, 2))
}

func sumArray(arr1 []int, arr2 []int) int {
	sum1 := 0
	sum2 := 0
	for i := 0; i < len(arr1); i++ {
		sum1 += arr1[i]
	}
	for j := 0; j < len(arr2); j++ {
		sum2 += arr2[j]
	}
	return sum1 + sum2

}

func mulArray(arr1 []int, arr2 []int) int {
	sum1 := 1
	sum2 := 1
	for i := 0; i < len(arr1); i++ {
		sum1 *= arr1[i]
	}
	for j := 0; j < len(arr2); j++ {
		sum2 *= arr2[j]
	}
	return sum1 * sum2

}

func reverseArray(arr1 []int) {
	fmt.Println(arr1)
	arr2 := make([]int, len(arr1))
	j := 0
	for i := len(arr1) - 1; i >= 0; i-- {
		arr2[j] = arr1[i]
		j++
	}
	fmt.Println("Origianl array", arr1, "Reversed Array", arr2)
	fmt.Println("Length of array", len(arr1))
	var sumArray int
	for i := 0; i < len(arr1); i++ {
		sumArray += arr1[i]
	}
	fmt.Println("Sum of array is ", sumArray)
	fmt.Println("First element of array is ", arr1[0], "Last element is ", arr1[len(arr1)-1])
}
