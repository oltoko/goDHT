/**
* goDHT a "Distributed Hash Table" library for the go language
*
* This file is part of goDHT.
*
* goDHT is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* goDHT is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with goDHT.  If not, see <http://www.gnu.org/licenses/>.
 */
package rTable

import (
	"goDHT"
	"goDHT/node"
	"math/big"
	"time"
)

var (
	CAddNode = make(chan *node.Node)
	CGetNode = make(chan GetNode)
	CStop    = make(chan int)
)

var (
	buckets []bucket
)

const (
	MAX_BUCKET_NODES = 8
)

func InitAndManage(storePath string) {

	store = storePath
	buckets = load()

	go manage()
}

func manage() {

	for {
		select {

		case node := <-CAddNode:

			addNode(node)

		case get := <-CGetNode:

			node := getNode(get.ID)

			if node != nil {
				get.Callback <- node
			}

			close(get.Callback)

		case <-CStop:

			save()
			break
		}
	}
}

type GetNode struct {
	Callback chan *node.Node
	ID       node.ID
}

type bucket struct {
	min         *big.Int
	max         *big.Int
	nodes       map[string]*node.Node // The Key is the ID as String
	lastChanged time.Time
}

func (b bucket) isFull() bool {

	if len(b.nodes) < MAX_BUCKET_NODES {
		return false
	} else {
		return true
	}
}

func (b bucket) match(id node.ID) bool {

	/* When a node with ID "N" is inserted into the table,
	it is placed within the bucket that has min <= N < max. */

	intID := id.Int()

	if b.min.Cmp(intID) > 0 {
		return false
	}

	if b.max.Cmp(intID) <= 0 {
		return false
	}

	return true
}

func (b bucket) addNode(n *node.Node) bool {

	if b.match(n.ID()) {
		b.nodes[n.ID().String()] = n
		return true
	}

	return false
}

func (b bucket) split() (bucket, bucket) {
	/* ... the bucket is replaced by two new buckets each with half the range of
	the old bucket and the nodes from the old bucket are distributed among the
	two new ones. */

	mid := big.NewInt(0).Div(b.max, big.NewInt(2))

	newB1 := bucket{
		min:         b.min,
		max:         mid,
		nodes:       make(map[string]*node.Node),
		lastChanged: time.Now(),
	}

	newB2 := bucket{
		min:         mid,
		max:         b.max,
		nodes:       make(map[string]*node.Node),
		lastChanged: time.Now(),
	}

	for _, n := range b.nodes {
		newB1.addNode(n)
		newB2.addNode(n)
	}

	return newB1, newB2
}

func addNode(n *node.Node) bool {

	splitted := false

	for i, b := range buckets {

		if !b.match(n.ID()) {
			continue
		}

		if !b.isFull() {
			b.addNode(n)
			return true
		}

		if b.match(goDHT.OwnNode().ID()) {
			/* When a bucket is full of known good nodes,
			no more nodes may be added unless our own node ID
			falls within the range of the bucket. */
			newB1, newB2 := b.split()
			buckets[i] = newB1
			buckets = append(buckets, newB2)
			splitted = true
		}
	}

	if splitted {
		return addNode(n)
	}

	return false
}

func getNode(id node.ID) *node.Node {

	var result *node.Node
	result = nil

	for _, b := range buckets {

		if b.match(id) {
			b.nodes
		}
	}

	return result
}
