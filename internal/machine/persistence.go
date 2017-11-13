package machine

import "example/internal/platform/persistence"

var persistor persistence.Persistor

func init() {
	persistor = &memoryPersistor{machines:map[int]Machine{}}
}

func Save(machine Machine) error {
	return persistor.Save(machine)
}

func Load(id int) (Machine, error) {
	entity, err := persistor.Load(id)
	machine := entity.(Machine)

	return machine, err
}

func Remove(id int) error {
	return persistor.Remove(id)
}

type memoryPersistor struct {
	machines map[int]Machine
}

func (m *memoryPersistor) Save(interface{}) error {
	return nil
}

func (m *memoryPersistor) Load(id int) (interface{}, error) {
	return Machine{}, nil
}

func (m *memoryPersistor) Remove(id int) error {
	return nil
}
