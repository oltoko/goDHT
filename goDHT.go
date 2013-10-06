// goDHT project goDHT.go
package goDHT

import (
	"goDHT/node"
)

var (
	ownNode = node.New(node.GenerateID())
)

func OwnNode() *node.Node {
	return ownNode
}
