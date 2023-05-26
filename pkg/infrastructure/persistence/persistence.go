package persistence

import (
	"terraform-lambda-dynamodb/pkg/domain/repository"

	"github.com/guregu/dynamo"
)

type Repositories struct {
	db             *dynamo.DB
	UserRepository repository.IUserRepository
}

func NewNewRepositories(db *dynamo.DB) *Repositories {
	return &Repositories{
		db:             db,
		UserRepository: NewUserRepository(db),
	}
}
