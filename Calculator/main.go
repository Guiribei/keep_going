package main

import "fmt"

func main() {
	var operation string
	var num1, num2 float32

	fmt.Println("Calculator GO 1.0:")
	fmt.Println("==================")
	
	fmt.Println("Which operation do you want to perform? (add, subtract, multiply, divide)")
	fmt.Scanf("%s", &operation)

	fmt.Println("Enter the first number: ")
	fmt.Scanf("%g", &num1)

	fmt.Println("Enter the second number: ")
	fmt.Scanf("%g", &num2)

	switch operation {
	case "add":
		fmt.Println(num1 + num2)
	case "subtract":
		fmt.Println(num1 - num2)
	case "multiply":
		fmt.Println(num1 * num2)
	case "divide":
			if num2 != 0 {
				fmt.Println(num1 / num2)
			} else {
				fmt.Println("Aí não né amigão!")
			}
	}

}