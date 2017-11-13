package persistence

type Persistor interface {
	Save(interface{}) error
	Load(id int) (interface{}, error)
	Remove(id int) error
}
