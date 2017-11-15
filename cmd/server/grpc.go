package server

import (
	"example/internal/machine"
	"example/internal/account"
)

type Grpc struct{}

func (*Grpc) HandleSomeRequest() error {
	//testing events when machine is updated
	err := machine.UpdateMachine()
	if err != nil {
		return err
	}

	//testing events when account is updated
	err = account.UpdateAccount()
	if err != nil {
		return err
	}

	return nil
}
