package main

import (
	"log"

	"github.com/5p4c351ck/DHT-InterNodal/node"
)

func main() {

	node, err := node.CreateNode()
	if err != nil {
		log.Fatal("Node creation failed")
		return
	}
	node.Store(node.Cid())
}
