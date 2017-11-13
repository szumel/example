package sync

func GetAccount(accountLogin string) (Account, error) {
	/*
		1. get account within package facade by login
		2. get relations (how and where store it?)
		3. get each machine within package facade (depending on relations)
		4. create sync.Machine for each machine
		4. create and return sync.Account with sync.Machines objects
	*/
}
