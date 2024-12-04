package main

import (
	"log"
	"time"

	"github.com/5p4c351ck/DHT-InterNodal/node"
)

func main() {

	node1, err := node.NewNode("127.0.0.1", "8080")
	if err != nil {
		log.Fatal("Remote node creation failed")
		return
	}

	node2, err := node.NewNode("127.0.0.1", "50001")
	if err != nil {
		log.Fatal("Remote node creation failed")
		return
	}

	nd1, err := node.NewLocalNode(node1)
	if err != nil {
		log.Fatal("Local node creation failed")
		return
	}

	nd2, err := node.NewLocalNode(node2)
	if err != nil {
		log.Fatal("Local node creation failed")
		return
	}

	go func() {
		nd1.Server()
	}()

	time.Sleep(2 * time.Second)
	nd2.Ping(nd1.Node)

}
