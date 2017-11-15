package account

import "example/events"

func UpdateAccount() error {
	events.DispatchChangedAccountLogin(events.MessageAccountLogin{NewLogin:"newLogin", OldLogin:"oldLogin"})

	return nil
}
