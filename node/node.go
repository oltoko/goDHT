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
	"fmt"
	"io"
	"math"
	"math/big"
	"time"
)

const (
	BAD_QUERY_DURATION time.Duration = time.Minute * 15
)

type ID []byte

func (id ID) Int() *big.Int {
	return big.NewInt(0).SetBytes(id)
}

func (id ID) String() string {
	return hex.EncodeToString(id)
}

func BiggestID() ID {

	biggestID := SmallestID()

	for i := range biggestID {
		biggestID[i] = math.MaxUint8
	}

	return biggestID
}

func SmallestID() ID {
	return make([]byte, 20)
}

type Node struct {
	id            ID
	lastQuery     time.Time
	lastResponse  time.Time
	onceResponded bool
}

func New(id ID) *Node {

	initTime := badTime()

	return &Node{
		id:            id,
		lastQuery:     initTime,
		lastResponse:  initTime,
		onceResponded: false,
	}
}

func (n Node) String() string {
	return fmt.Sprintf("{ id: %s, lastQuery: %s, lastResponse: %s, onceResponded: %t }",
		n.id.String(),
		n.lastQuery.String(),
		n.lastResponse.String(),
		n.onceResponded)
}

/*
A good node is a node has responded to one of our queries within the last
15 minutes.
A node is also good if it has ever responded to one of our queries
and has sent us a query within the last 15 minutes.
*/
func (n Node) IsGood() bool {

	if n.lastResponse.After(badTime()) {
		return true
	} else if n.lastQuery.After(badTime()) && n.onceResponded {
		return true
	} else {
		return false
	}

}

/*
In Kademlia, the distance metric is XOR and the result is interpreted as an unsigned integer.
distance(A,B) = |A xor B| Smaller values are closer.
*/
func (n Node) Distance(id ID) *big.Int {
	return big.NewInt(0).Xor(n.id.Int(), id.Int())
}

func (n Node) ID() ID {
	return n.id
}

func (n Node) LastQuery() time.Time {
	return n.lastQuery
}

func (n *Node) SetLastQuery() {
	n.lastQuery = time.Now()
}

func (n Node) LastResponse() time.Time {
	return n.lastResponse
}

func (n *Node) SetLastResponse() {
	n.lastResponse = time.Now()
}

func (n Node) HasOnceResponded() bool {
	return n.onceResponded
}

func (n *Node) SetOnceResponded() {
	n.onceResponded = true
}

/*
A Bad Time is BAD_QUERY_DURATION before time.Now()
*/
func badTime() time.Time {
	return time.Now().Add(-BAD_QUERY_DURATION)
}

func GenerateID() ID {

	hash := sha1.New()

	p := make([]byte, 4096)
	io.ReadFull(rand.Reader, p)

	hash.Write(p)

	return hash.Sum(nil)
}
