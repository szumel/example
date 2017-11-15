package server

import (
	"example/internal/machine"
)

type Grpc struct{}

func (*Grpc) HandleSomeRequest() error {
	err := machine.UpdateMachine()
	if err != nil {
		return err
	}

	return nil
}
