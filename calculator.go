package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// Calculator function to perform operations based on the operand
func Calculator(value1, value2 float64, operand string) (float64, error) {
	// TODO: Implement Calculator function
	// implement the switch inside the functions
	// Add all the basic operands using switch
	switch calculator{
	case "+" : 
	return value1 + value2,nil
	case "_" :
		return value1 - value2,nil
	case "*" :
		return value1 * value2,nil
	case "\" :
		if (value2 == 0){
			return 0,nil}

			return value1 / value2,nil
		
	}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression (e.g., 10 + 20): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	// TODO: Check if exact 3 parts are given. If not, throw an error
	if len(s.elements)!=3{
		fmt.Println("Invalid")
		return
	}
	// TODO: Take all 3 values and pass to calculator function
	// Pass the value1 ,value2, and operand to the calculator funtion
	value1,err:= strconv.ParseFloat(parts[0],64)
	if err!= nil{
		fmt.Println("Invalid operators")
		return
	}
	operand := parts[1]
	value2,err := strconv.ParseFloat(parts[2])
	if err!=nil{
		fmt.Println("Invalid operators")
		return
	}
	
	// TODO: Print results
	result = Calculator(value1,value2,operand)
	if err!=nil{
		return 0,nil
	}
    
	fmt.Printf("Result: %v\n", result)
}

















