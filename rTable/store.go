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
	"goDHT/node"
	"time"
)

var (
	store string
)

func load() []bucket {

	if storeExist() {
		// ...
		return nil
	}

	buckets := make([]bucket, 1)

	// An empty table has one bucket with an ID space range of min=0, max=2^160
	buckets[0] = bucket{
		min:         node.SmallestID().Int(),
		max:         node.BiggestID().Int(),
		nodes:       make(map[string]*node.Node),
		lastChanged: time.Now(),
	}

	return buckets
}

func storeExist() bool {
	return false
}

func save() error {
	return nil
}
