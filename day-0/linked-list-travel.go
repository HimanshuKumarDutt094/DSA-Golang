// package main

// import "fmt"

// // linkedlist is is cllection of nodes each having pair of data and pointer to next node , its memory efficent becuase
// // it works with non contiguous manner, i.e., if nextblock is occupied ,ittl check nect free block and link it for next element
// type Node struct {
// 	data     int
// 	nextNode *Node
// }

// func main() {
// 	linkedList := make([]Node, 4) //makes array of nodes each node is a memory block having
// 	//data and a pointer which points to next memory block
// 	fmt.Println("enter a digit: ") //user input
// 	var v int                      //temp variable to store input
// 	for i := 0; i < len(linkedList); i++ {
// 		_, err := fmt.Scan(&v) //scan user input and assign it to v
// 		if err != nil {        //if error while scanning print it
// 			fmt.Print(err)
// 		}
// 		linkedList[i].data = v //assing user input to node.data

// 		//following [0] make the last node point to first node making it circlular linkedlist
// 		//or nil to make it normal linked list
// 		if i == len(linkedList)-1 {

// 			linkedList[i].nextNode = nil
// 			continue
// 		}
// 		linkedList[i].nextNode = &linkedList[i+1] //assign the pointer address of next node & [i+1]

// 	}
// 	c := &linkedList[0] //first node get O(1)
// 	i := 0              //temp var in case circlular loop
// 	fmt.Println("printing linked list:")
// 	for c != nil { //check of next node is nil or not, nil inicated end of linked list
// 		fmt.Println("data is :", c.data) //print current node data
// 		c = c.nextNode                   //set c to next node in memory block
// 		if i == 15 {                     // limit the loop to 15 in case circluar linked list is found
// 			break
// 		}
// 		i++ //increment i , in cae of normal loop will stop at 10
// 	}
// 	//or
// 	fmt.Println(linkedList)
// }
