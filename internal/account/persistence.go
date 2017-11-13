package account

import "example/internal/platform/persistence"

var persistor persistence.Persistor

func init() {
	persistor = &memoryPersistor{accounts:map[int]Account{}}
}

func Save(account Account) error {
	return persistor.Save(account)
}

func Load(id int) (Account, error) {
	entity, err := persistor.Load(id)
	account := entity.(Account)

	return account, err
}

func Remove(id int) error {
	return persistor.Remove(id)
}

type memoryPersistor struct {
	accounts map[int]Account
}

func (m *memoryPersistor) Save(interface{}) error {
	return nil
}

func (m *memoryPersistor) Load(id int) (interface{}, error) {
	return Account{}, nil
}

func (m *memoryPersistor) Remove(id int) error {
	return nil
}
