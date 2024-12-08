package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/5p4c351ck/InterNodal/controller"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	contrl, err := controller.NewController()
	if err != nil {
		return
	}
	for {
		fmt.Printf(">> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		contrl.Receive(input)
		contrl.Handle()
	}
}
