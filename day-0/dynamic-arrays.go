// package main

// import "fmt"

// func main() {
// 	//dynamic arrays are arrays which can be resized anytime when limit is reached (cap)
// 	// in golang these are called slices (similar to python)
// 	slice1 := []int{1, 2} //initial slice of cap 2 len 2 but can be modified on whim
// 	// or
// 	// slice2 = make([]int, 10)

// 	fmt.Println("slice 1: ", slice1)
// 	fmt.Println("slice 1 cap: ", cap(slice1))
// 	fmt.Println("slice 1 len: ", len(slice1))
// 	//if we do soemthing like
// 	//	slice1[6] = 9
// 	//index out of range but if we do following
// 	slice1 = append(slice1, 1, 2, 3, 4, 5, 6, 7, 8)
// 	//since slice is full go will create a new slice with x2 cap of pev slice and append all elements to new slice
// 	//lookup is always constant
// 	fmt.Println("slice 1: ", slice1)
// 	fmt.Println("slice 1 cap: ", cap(slice1))
// 	fmt.Println("slice 1 len: ", len(slice1))
// 	//10 10
// 	//looping is same
// 	fmt.Println(slice1[:4])
// 	fmt.Println(slice1[5:])
// 	fmt.Println(slice1[len(slice1)-1:]) //
// 	fmt.Println(slice1[5 : len(slice1)-1])
// 	// -ve indexing is not supported fmt.Println(slice1[-1:4])

// }
