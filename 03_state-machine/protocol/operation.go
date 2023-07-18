package protocol

import (
	"fmt"
	"strings"
)

// Operator is a type of operation applied to the state machine
type Operator string

const (
	command_add = "+" // Operator add
	command_sub = "-"
	command_mul = "*"
	command_div = "/"
)

// Translate generates an operator from a string
func Translate(o string) (Operator, error) {
	for k, v := range map[string]string{command_add: "add", command_sub: "sub", command_mul: "mul", command_div: "div"} {
		if strings.EqualFold(o, v) {
			return Operator(k), nil
		}
	}
	return Operator("UNKNOW"), fmt.Errorf("Unknown operator: %s", o)
}

// Apply applies the operation to the inputs, and returns the result of the operation
func (o *Operator) Apply(lhs, rhs int) (int, error) {
	switch *o {
	case command_add:
		return lhs + rhs, nil
	case command_sub:
		return lhs - rhs, nil
	case command_mul:
		return lhs * rhs, nil
	case command_div:
		if rhs == 0 {
			return 0, fmt.Errorf("Zero Division")
		}
		return lhs / rhs, nil
	}
	return 0, fmt.Errorf("Unknown operator: %s", *o)
}

// String returns the operation in the form of string
func (o *Operator) String() string {
	return string(*o)
}

// RequestCommandArgs notifies the operation with an operand to the server
type RequestCommandArgs struct {
	Operator Operator
	Operand  int
}

// RequestCommandReply returns the result of the operation to the client
type RequestCommandReply struct {
	Value int
	Error error
}
