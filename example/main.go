package main

import (
	"fmt"
	"goDHT/node"
)

func main() {

	id := node.GenerateID()

	node1 := node.New(id)
	node2 := node.New(id)

	fmt.Println(node1.Distance(node2.ID()))

	fmt.Println("--------------")

	biggestID := node.BiggestNodeID()
	fmt.Println(biggestID.String())
	fmt.Println(biggestID.Int().String())

	fmt.Println("--------------")

	for i := 0; i < 10; i++ {
		fmt.Println(node.GenerateID())
	}
}
