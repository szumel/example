package account

import (
	"example/internal/platform/persistence"
)

type Account struct {
	Login string
	persistence.Meta
}
