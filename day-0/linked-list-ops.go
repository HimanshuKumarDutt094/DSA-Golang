package main

import "fmt"

// operations in linked list
// traverse, insert, delete, search , exist
type Node struct {
	data int
	next *Node
}
type LinkedList struct { //create a linked list (it starts with head node)
	head *Node
}

type LinkedListInterface interface { //declate function related to linkedList via interface
	//(class functions in other languages , go only has struct and interface to achieve oop)
	traverse()
	insertHead(value int)
	insertEnd(value int)
	insertAny(value int, position int)
	deleteHead()
	deleteEnd()
	deleteAnyByPosition()
	deleteAnyByValue()
	exist(value int) bool
}

// the () after func is called function signature , it links or relates the written function to whatever struct (class) we wants
// note we need to do this so we can call LinkedList.traverse()
// the memory address of LinkedList is given so that theres only one linkedList with and we can manipulate it directly
func (linkedlist *LinkedList) traverse() {
	current := linkedlist.head //set the current var to head of linkedList it has data and next
	if current == nil {
		fmt.Println("empty list")
	}
	for current != nil { //if the head is null skips else loop over it
		if current.next == nil { // if the current node .next is nil mean we are at end so print the data
			fmt.Printf("%d", current.data)

		} else {
			//else print -> to indicate linkdlist
			fmt.Printf("%d ->", current.data)
		}
		current = current.next //move the index to next node (.next points to node itself not data or something)
	}
	fmt.Println()
}
func (linkedlist *LinkedList) insertHead(value int) { //take user value
	t := &Node{data: value}  //create a new node pointer with data as user value
	t.next = linkedlist.head //set the temp node next to head of linked list
	linkedlist.head = t      //set the head pointer to the temp node
	// t.next,linkedList.head=linkedList.head,t can work as well a,b=b,c ,sets -> a=b b=c
	fmt.Println("head inserted successfully: ", linkedlist)
}
func (linkedlist *LinkedList) insertEnd(value int) {
	t := &Node{data: value} // create the new node

	// If the list is empty, set the new node as the head
	if linkedlist.head == nil {
		linkedlist.head = t
		return
	}

	// Traverse to the last node
	c := linkedlist.head
	for c.next != nil {
		c = c.next
	}

	// Set the new node as the last node's next
	c.next = t
	fmt.Println("tail inserted successfully: ", linkedlist)

}
func (linkedlist *LinkedList) insertAny(value int, position int) { //value and position
	current := linkedlist.head //start from head
	t := &Node{data: value}    //new temp node
	if position == 1 {
		linkedlist.insertHead(value)
		return
	}
	if current == nil {
		current = t
		return
	}
	i := 1               //we are using 1 based indexing
	for current != nil { //loop over linked list
		if i == position-1 { //1 index before the desired position
			// set the temp node next value to current node next (why?)
			t.next = current.next // in 6->4->1->0 position is 3 (data =1) we stop at 4 and set next of 4 to next of temp
			//6->4-->(1)
			//(temp)-^

			current.next = t //we also set current .next or next value t ne the new node address
			//basically below diagram
			//6->4-      1- >0
			//     |     ^
			//     V     |
			//    (temp)-I
			break
		}
		//loop condition
		i++
		current = current.next
	}
	fmt.Printf(" inserted successfully at %v ", position)
	fmt.Println()

}
func (linkedlist *LinkedList) deleteHead() {
	if linkedlist.head == nil {
		return
	}
	linkedlist.head = linkedlist.head.next
}
func (linkedlist *LinkedList) deleteEnd() {
	current := linkedlist.head
	if current == nil {
		fmt.Println("empty list")
	}
	for current != nil {
		if current.next.next == nil {
			current.next = nil
			return
		}
		current = current.next
	}

}
func (linkedlist *LinkedList) deleteAnyByPosition(position int) {
	i := 1
	current := linkedlist.head
	if current == nil {
		fmt.Println("list empty")
	}
	for current != nil {
		if i+1 == position {
			fmt.Println("removed ", current.next.data)
			current.next = current.next.next //set current node next = to next of next of current basically skip 1 node or the node we want to remove
			return
		}
		i++
		current = current.next
	}

}
func (linkedlist *LinkedList) deleteAnyByValue(value int) {
	i := 1
	current := linkedlist.head
	if current == nil {
		fmt.Println("list empty")
	}
	for current != nil {
		if current.next.data == value {
			fmt.Printf("\nremoved %d position found at %d\n", current.next.data, i+1)

			current.next = current.next.next //set current node next = to next of next of current basically skip 1 node or the node we want to remove
			return
		}
		i++
		current = current.next
	}

}
func (linkedlist *LinkedList) exist(value int) bool {
	i := 1
	current := linkedlist.head
	for current != nil {
		if current.data == value {

			fmt.Println("found at index :", i)
			return true
		}
		current = current.next
		i++
	}
	return false

}
func main() {
	linkedList := LinkedList{}
	linkedList.traverse()

	for i := 0; i < 2; i++ {
		linkedList.insertHead(i)
	}

	linkedList.traverse()
	linkedList.insertAny(50, 1)
	linkedList.traverse()
	linkedList.deleteHead()
	linkedList.traverse()
	linkedList.deleteEnd()
	linkedList.traverse()
	linkedList.insertAny(53, 1)
	linkedList.insertAny(0, 1)
	linkedList.insertAny(33, 1)
	linkedList.insertAny(20, 1)
	linkedList.traverse()

	linkedList.deleteAnyByPosition(3)
	linkedList.traverse()

	linkedList.deleteAnyByValue(53)
	linkedList.traverse()
	linkedList.exist(1)
}
