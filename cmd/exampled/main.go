package main

import "example/cmd/server"

func main() {
	//TCP listen etc.
	s := server.Grpc{}
	s.HandleSomeRequest()
}
