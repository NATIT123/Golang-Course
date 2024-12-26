package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	///Ex1
	ex1()
	//Ex2
	fmt.Println("Result of Ex2 is:", ex2())
	//Ex3
	ex3()

	//Ex4
	fmt.Println("Result of Ex4", twoSum([]int{2, 7, 11, 15}, 9))
}

//Ex1
func ex1() {
	fmt.Println("Ex1")
	var a, b int
	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)
	area := a * b
	fmt.Println("Area of rectangle is:", area)
	premiter := 2 * (a + b)
	fmt.Println("Perimeter of rectangle is:", premiter)
}

//Ex2
func ex2() bool {
	fmt.Println("Ex2")
	var a string
	fmt.Println("Enter a string:")
	fmt.Scan(&a)
	lengtofstring := len(a)
	return lengtofstring%2 == 0
}

//Ex3
func ex3() {
	fmt.Println("Ex3")
	var n int
	fmt.Println("Enter length of array")
	fmt.Scan(&n)

	//Assume capacity of array is 100
	var arr [100]int
	for i := 0; i < n; i++ {
		fmt.Println("Enter element", i)
		fmt.Scan(&arr[i])
	}

	sliceNumber := arr[0:n]
	//Total of slice
	total := 0
	for i := 0; i < len(sliceNumber); i++ {
		total += sliceNumber[i]
	}
	fmt.Println("Total of slice is:", total)

	//Min and max of slice
	minNumber := sliceNumber[0]
	fmt.Println("Min of slice is:", minNumber)
	maxNumber := sliceNumber[0]
	for i := 1; i < len(sliceNumber); i++ {
		if minNumber > sliceNumber[i] {
			minNumber = sliceNumber[i]
		}
		if maxNumber < sliceNumber[i] {
			maxNumber = sliceNumber[i]
		}
	}

	fmt.Println("Min of slice is:", minNumber)
	fmt.Println("Max of slice is:", maxNumber)

	//Sort slice
	for i := 0; i < len(sliceNumber); i++ {
		for j := i + 1; j < len(sliceNumber); j++ {
			if sliceNumber[i] > sliceNumber[j] {
				temp := sliceNumber[i]
				sliceNumber[i] = sliceNumber[j]
				sliceNumber[j] = temp
			}
		}
	}

	fmt.Println("Sorted slice is:", sliceNumber)
}

//Ex4
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		value, exists := hashMap[target-num]
		if exists {
			return []int{value, i}
		}
		hashMap[num] = i
	}
	return []int{}
}
