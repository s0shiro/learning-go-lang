package _switch

import (
	"fmt"
	"time"
)

func SwitchStatement() {
	i := 2
	message := fmt.Sprintf("Write %d as ", i)
	fmt.Print(message)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ahh yes its weekend hooray!")
	default:
		fmt.Println("Ahh sad its weekday.")
	}

	//switch as alternate to else/if statement
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon!")
	default:
		fmt.Println("its after noon!")
	}

	//A type switch compares types instead of values.
	//You can use this to discover the type of an interface value.
	//In this example, the variable t will have the type corresponding to its clause.

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Type is boolean.")
		case int:
			fmt.Println("Type is integer.")
		default:
			fmt.Println("Unknown type.", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")

}
