package server

import (
	"example/cmd/sync"
	"fmt"
)

type Grpc struct{}

func (*Grpc) HandleSomeRequest() error {
	/*
	suppose magically request handling here
	 */

	acc, err := sync.GetAccount("example@example.com")
	if err != nil {
		return err
	}

	//perform actions with account
	fmt.Println(acc)

	return nil
}
