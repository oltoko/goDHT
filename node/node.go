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
package node

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"math"
	"math/big"
	"time"
)

const (
	BAD_QUERY_DURATION time.Duration = time.Minute * 15
)

type NodeID []byte

func (id NodeID) Int() *big.Int {
	return big.NewInt(0).SetBytes(id)
}

func (id NodeID) String() string {
	return hex.EncodeToString(id)
}

func BiggestNodeID() NodeID {

	biggestID := SmallestNodeID()

	for i := range biggestID {
		biggestID[i] = math.MaxUint8
	}

	return biggestID
}

func SmallestNodeID() NodeID {
	return make([]byte, 20)
}

type Node struct {
	id            NodeID
	lastQuery     time.Time
	lastResponse  time.Time
	onceResponded bool
}

func New(id NodeID) *Node {

	initTime := badTime()

	return &Node{
		id:            id,
		lastQuery:     initTime,
		lastResponse:  initTime,
		onceResponded: false,
	}
}

func (n Node) IsGood() bool {

	/* A good node is a node has responded to one of our queries within the last
	15 minutes. A node is also good if it has ever responded to one of our queries
	and has sent us a query within the last 15 minutes.*/
	if n.lastResponse.After(badTime()) {
		return true
	} else if n.lastQuery.After(badTime()) && n.onceResponded {
		return true
	} else {
		return false
	}

}

func (n Node) Distance(id NodeID) *big.Int {
	return big.NewInt(0).Xor(n.id.Int(), id.Int())
}

func (n Node) ID() NodeID {
	return n.id
}

func (n Node) LastQuery() time.Time {
	return n.lastQuery
}

func (n Node) LastResponse() time.Time {
	return n.lastResponse
}

func (n Node) HasOnceResponded() bool {
	return n.onceResponded
}

func (n Node) SetOnceResponded() {
	n.onceResponded = true
}

func badTime() time.Time {
	return time.Now().Add(-BAD_QUERY_DURATION)
}

func GenerateID() NodeID {

	hash := sha1.New()

	p := make([]byte, 4096)
	io.ReadFull(rand.Reader, p)

	hash.Write(p)

	return hash.Sum(nil)
}
