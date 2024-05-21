// Calculate the sum of an array and return the average
package main

import (
	"fmt"
)

func calAvg(slice []int) int {
	count := 0
	for i := 0; i < len(slice); i++ {
		count += slice[i]
	}
	return count / len(slice)
}

// Function to calculate age

func checkAge(age int) {

	if age < 0 {
		fmt.Println("Invalid Entry")
	} else if age > 0 && age < 18 {
		fmt.Println("Minor")
	} else if age >= 18 && age <= 25 {
		fmt.Println("Young Adult")
	} else if age >= 25 && age <= 120 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Invalid option")
	}

}

// Reverse the string
func reverseString(s string) string {
	// get the length
	length := len(s)
	//hold the value of reversed
	reversed := make([]byte, length)
	for i := 0; i < length; i++ {
		reversed[i] = s[length-1-i]
	}
	return string(reversed)
}

// Largest Number
func largestNumber(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// counter
type counter struct {
	c int
}

func increment(count counter) counter {
	count.c++
	return count
}
func decrement(count counter) counter {
	count.c--
	return count
}

func main() {
	slice := []int{10, 20, 50, 10, 10}
	fmt.Println(calAvg(slice))
	// fmt.Println(checkAge(age))
	checkAge(23)
	str := "hello"
	reversed := reverseString(str)
	fmt.Println(reversed)
	arr := []int{23, 556, 233}
	fmt.Println((largestNumber(arr)))

	p := counter{}
	p1 := increment(increment(p))
	fmt.Printf("%d", p1)
	p2 := decrement(p1)
	fmt.Printf("%d", p2)
	p3 := decrement(p2)
	fmt.Printf("%d", p3)
}
