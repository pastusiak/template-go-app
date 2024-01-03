package console

import (
	"fmt"
)

type Example struct{}

func (e *Example) CmdConfig() (trigger string, description string) {
	trigger = "e1"
	description = "Example description"
	return
}

func (e *Example) CmdExecute() bool {
	fmt.Println("Lorem ipsum dolar")
	return true
}
