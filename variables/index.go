package variables

import "fmt"

const s string = "constant"

func Variables() {
	var num1, num2 int = 1, 2
	fmt.Println(num1 + num2)

	//shorthand
	name := "Niel"
	age := 21
	isPogi := true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Pogi:", isPogi)
}

func Constant() {
	fmt.Println(s)
}
