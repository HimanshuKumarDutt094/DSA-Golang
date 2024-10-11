// package main

// import "fmt"

// // in go array / list can be static or dynamic.
// // static arrays have fixed memory
// // it can be initialized via following
// // note array can be of anything, int ,runes (integer representation of chars), strings , bools , floats
// // we will go with int (it can be 32 64 depending on user system)
// func main() {
// 	array1 := make([]int, 10, 10) //using make to create array with capacity and length (hover over to see a suggestion)
// 	//cap is max capacity array can hold
// 	//len is current length of array, i can be increased to match cap
// 	fmt.Println("array 1", array1)
// 	fmt.Println("capacity of array 1", cap(array1))
// 	fmt.Println("length of array 1", len(array1))
// 	array2 := [10]int{}
// 	fmt.Println("array 2", array2)
// 	fmt.Println("capacity of array 2", cap(array2)) //output 10 because : by default cap and len = size given to static array
// 	fmt.Println("length of array 2", len(array2))   //output 10
// 	//use array[index]= value to modify or
// 	//array=append(array,value1) to appned to array (more useful in slices)
// fmt.Println(array1[:4])
// fmt.Println(array1[5:])
// fmt.Println(array1[len(array1)-1:]) //
// fmt.Println(array1[5 : len(array1)-1])
// 	rangeOverStaticArrayViaForLoop([]int{1, 2, 3})
// }
// func rangeOverStaticArrayViaForLoop(array []int) {
// 	// there are various ways to range over something
// 	// basic for loop
// 	for i := 0; i < len(array); i++ {
// 		fmt.Println(array[i])
// 	}
// 	// reverse
// 	for i := len(array) - 1; /*len give size but we follow 0 index so -1 the len*/ i >= 0; i-- {
// 		fmt.Println(array[i])
// 	}
// 	// go doesnt have while so we can use for in following way
// 	var i int = 0
// 	for i < len(array) {
// 		fmt.Println(array[i])
// 		i++
// 	}
// 	//or like this
// 	i = len(array) - 1
// 	for i >= 0 {
// 		fmt.Println(array[i])
// 		i--
// 	}
// 	// in go theres also range which makes stuff easier but , range doesnt gurantee order in case of hash array sometimes
// 	for index, value := range array {
// 		fmt.Printf("\nindex is %v, value is %v \t", index, value)
// 	}
// }
