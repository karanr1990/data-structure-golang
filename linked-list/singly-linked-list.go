package main

import (
	"fmt"
	"math/rand"
)

//create node struct
type Node struct {
	data int
	next *Node
}

//create node list struct
type NodeList struct {
	length int
	nodes  *Node
}

func (nodeList *NodeList) Display() {

	lst := nodeList.nodes

	for lst != nil {
		fmt.Printf("%v -> ", lst.data)
		lst = lst.next
	}
	fmt.Println()

}

func (lst *NodeList) insertNode(newNode *Node) {

	if lst.length == 0 {
		lst.nodes = newNode
	} else {
		currentNode := lst.nodes

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = newNode
	}

	lst.length++

}

func main() {

	items := &NodeList{}
	size := 10

	for i := 0; i < size; i++ {

		node := Node{data: rand.Intn(100)}

		if node.data == 0 {
			i = i - 1
			continue
		}

		items.insertNode(&node)

		fmt.Printf("%v Length: %v and number is %v\n", i, items.length, node.data)
	}

	items.Display()

}
