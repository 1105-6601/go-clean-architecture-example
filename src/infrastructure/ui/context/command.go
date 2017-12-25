package context

import (
	"fmt"
)

type Command struct {
	Arguments []interface{}
}

func NewCommand(args ...interface{}) *Command {
	return &Command{
		Arguments: args,
	}
}

func (cmd *Command) Echo(str string) {
	fmt.Println(str)
}

func (cmd *Command) Args() ([]interface{}) {
	return cmd.Arguments
}
