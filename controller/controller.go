package controller

import (
	"log"
	"strings"

	"github.com/5p4c351ck/DHT-InterNodal/dht"
)

const maxInputLength = 80

type Controller interface {
	Receive(input string)
	Handle()
}

type controllerImpl struct {
	command string
	DHT     dht.DHT
}

func (controller *controllerImpl) Receive(input string) {
	input = strings.TrimSpace(input)
	controller.command = input[:maxInputLength]
}

func (controller *controllerImpl) Handle() {
	switch controller.command {
	case "init":
		if controller.DHT != nil {
			log.Println("A DHT is already initialized")
			return
		}
		d, err := dht.NewDHT("0.0.0.0", "8080")
		if err != nil {
			return
		}
		controller.DHT = d
		log.Println("DHT initialized")
	default:
		log.Printf("%s is not a valid command", controller.command)
	}
}

func NewController() (Controller, error) {
	controller := &controllerImpl{}
	return controller, nil
}
