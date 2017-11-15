package events

import (
	"fmt"
	"example/events"
)

type Dispatcher struct {}

func (Dispatcher) ChangedName(name events.MessageChangedName) {
	fmt.Println("dispatching changedname from account")
}

