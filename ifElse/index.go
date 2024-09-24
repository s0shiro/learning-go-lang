package ifElse

import "fmt"

func ConditionalStatement() {
	//basic example
	if 7%2 == 0 {
		fmt.Println("7 is a even number")
	} else {
		fmt.Println("7 is odd number.")
	}

	//if statement w/out else statement
	if 6%3 == 0 {
		fmt.Println("6 is divisivible by 3.")
	}

	//using logical operators
	if 8%4 == 0 || 7%3 == 0 {
		fmt.Println("either  or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digit")
	}
}
