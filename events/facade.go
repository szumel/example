package events

func DispatchChangedName(message MessageChangedName) error {
	for _, dispatcher := range pool["changedName"] {
		dispatcher.ChangedName(message)
	}

	return nil
}