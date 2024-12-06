package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("---InterNodal---")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">> ")
		_, err := reader.ReadString('\n')
		if err != nil {
			return
		}
	}
}
