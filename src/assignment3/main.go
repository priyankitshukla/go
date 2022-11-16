package main

import "fmt"

func main() {

	// numbers := getArray()
	// findAllOddAndEvenNumber(numbers)
	// searchElementInArray(numbers)
	// sumAndmulArray(numbers)
	// sliceAndSort(numbers)
	createMapforEmp()

}

func createMapforEmp() {
	var empMap = make(map[int]string)
	var empId int
	var empName string
	for true {
		fmt.Println("Please enter Emp Id - ")
		fmt.Scan(&empId)
		fmt.Println("Please enter Emp Name - ")
		fmt.Scan(&empName)
		if int(empId) == -999 {
			break
		} else {
			empMap[empId] = empName
		}

	}
	fmt.Println(empMap)
}

func sliceAndSort(numbers []int) {
	var startIndex int
	var endIndex int
	fmt.Println("Please enter the start index:\t")
	fmt.Scan(&startIndex)
	fmt.Println("Please enter the end index:\t")
	fmt.Scan(&endIndex)
	X := numbers[startIndex:endIndex]
	for i := 0; i < len(X); i++ {
		for j := 0; j < len(X)-1; j++ {
			if X[j] > X[j+1] { //2,3
				X[j] = X[j] + X[j+1]   //5,3
				X[j+1] = X[j] - X[j+1] // 5,2
				X[j] = X[j] - X[j+1]

			}
		}
	}
	fmt.Println("sorted array ", X)
}

func sumAndmulArray(numbers []int) {
	sum, mul := 0, 1
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
		mul *= numbers[i]
	}
	fmt.Printf("Sum of array is %d and Multiplication of array is %d \n", sum, mul)
}

func searchElementInArray(numbers []int) {
	fmt.Println("Please enter the number you want to search")
	var num int
	fmt.Scan(&num)
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == num {
			fmt.Printf("Found element %d at index %d \n", num, i)
			break
		} else if i == len(numbers)-1 {
			fmt.Println("Element not found.")
		}
	}

}

func getArray() []int {
	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	var numbers = make([]int, size)

	for i := 0; i <= size-1; i++ {
		fmt.Printf("Please enter %d value \n", i+1)
		fmt.Scan(&numbers[i])
	}
	return numbers
}

func findAllOddAndEvenNumber(numbers []int) {
	odds := make([]int, len(numbers))
	evens := make([]int, len(numbers))
	oddCounter, evenCounter := 0, 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			evens[evenCounter] = numbers[i]
			evenCounter++
		} else {
			odds[oddCounter] = numbers[i]
			oddCounter++
		}
	}
	fmt.Println("Array of even number -->", evens)
	fmt.Println("Array of odd number -->", odds)

}
