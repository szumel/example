package events

var pool = map[string][]Dispatcher{}

type MessageChangedName struct {
	OldName string
	NewName string
}

type MessageAccountLogin struct {
	OldLogin string
	NewLogin string
}

type Dispatcher interface {
	ChangedMachineName(message MessageChangedName)
	ChangedAccountLogin(message MessageAccountLogin)
}

func RegisterDispatcher(eventCode string, dispatcher Dispatcher) {
	pool[eventCode] = append(pool[eventCode], dispatcher)
}


