package controller

const maxInputLength = 80

type Controller interface {
	Receive(input string)
	Handle()
}

type controllerImpl struct {
	data string
}

func (controller *controllerImpl) Receive(input string) {
	controller.data = input[:maxInputLength]
}

func (controller *controllerImpl) Handle() {}

func NewController() (Controller, error) {
	controller := &controllerImpl{}
	return controller, nil
}
