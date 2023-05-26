package repository

import (
	"github.com/guregu/dynamo"
)

type IUserRepository interface {
	Example(db *dynamo.DB)
}
