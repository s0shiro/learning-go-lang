package arrays

import "fmt"

func Arrays() {
	//initialize an array
	var arr [5]int
	fmt.Println("emp: ", arr)

	//add values on certain index of an array
	arr[4] = 100
	fmt.Println("set: ", arr)
	fmt.Println("get: ", arr[4])

	//len function return the length of an array
	fmt.Println("length: ", len(arr))

	//declare array in one line
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dec: ", arr2)

	//let the compiler count the number of elements w/ ... like spread in js
	arr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", arr2)

	//seems confusing but if you specify index with : the indexes between will become 0
	arr2 = [...]int{100, 3: 300, 500}
	fmt.Println("indx: ", arr2)

	//2 or multiple dimension of array
	var twoD [2][3]int

	//very tricky HAHAHAHAHA
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D Array: ", twoD)

	//You can create and initialize multi-dimensional arrays at once too.
	twoD = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println("2D: ", twoD)
}
