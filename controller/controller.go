package controller

import (
	"strings"
)

const maxInputLength = 80

type Controller interface {
	Receive(input string)
	Handle()
}

type controllerImpl struct {
	data string
}

func (controller *controllerImpl) Receive(input string) {
	input = strings.TrimSpace(input)
	controller.data = input[:maxInputLength]
}

func (controller *controllerImpl) Handle() {
	switch controller.data {
	case "init":
	}
}

func NewController() (Controller, error) {
	controller := &controllerImpl{}
	return controller, nil
}
