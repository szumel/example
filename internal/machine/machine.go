package machine

import "example/internal/platform/persistence"

type Machine struct {
	Name string
	persistence.Meta
}
