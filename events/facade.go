package events

func DispatchChangedMachineName(message MessageChangedName) error {
	for _, dispatcher := range pool["changedMachineName"] {
		dispatcher.ChangedMachineName(message)
	}

	return nil
}

func DispatchChangedAccountLogin(message MessageAccountLogin) error {
	for _, dispatcher := range pool["changedAccountLogin"] {
		dispatcher.ChangedAccountLogin(message)
	}

	return nil
}