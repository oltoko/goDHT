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
