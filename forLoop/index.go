package forLoop

import "fmt"

func ForLoop() {
	//basic for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	fmt.Println("---------------------------")
	//classic for loop
	for j := 0; j <= 3; j++ {
		fmt.Println(j)
	}

	//another way but not working btw GHAHAHAHA
	for i := range 4 {
		fmt.Println("range", i)
	}

	//use "break" or "return" to break the loop
	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
