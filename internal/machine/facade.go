package machine

import "example/events"

func UpdateMachine() error {
	/*
	executing some persistence here and dispatching event
	 */

	 err := events.DispatchChangedMachineName(events.MessageChangedName{OldName:"oldName", NewName:"newName"})
	 if err != nil {
	 	return err
	 }

	return nil
}