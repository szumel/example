//This package is responsible for relation between entities
package sync

import (
	"example/internal/account"
	"example/internal/machine"
)

type Machine struct {
	Accounts []account.Account
	machine.Machine
}

type Account struct {
	Machines []machine.Machine
	account.Account
}