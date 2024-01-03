package tasks

import "fmt"

type Example struct{}

func (t *Example) CmdConfig() (trigger string, description string) {
	trigger = "e3"
	description = "Example description"
	return
}

func (t *Example) CmdExecute() bool {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}

	return true
}
