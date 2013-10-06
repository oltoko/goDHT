// store
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
		min:         node.SmallestNodeID().Int(),
		max:         node.BiggestNodeID().Int(),
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
