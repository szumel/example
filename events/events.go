package events

var pool = map[string][]Dispatcher{}

type MessageChangedName struct {
	OldName string
	NewName string
}

type Dispatcher interface {
	ChangedName(message MessageChangedName)
}

func RegisterDispatcher(eventCode string, dispatcher Dispatcher) {
	pool[eventCode] = append(pool[eventCode], dispatcher)
}


