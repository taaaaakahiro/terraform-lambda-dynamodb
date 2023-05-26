package persistence

import (
	"terraform-lambda-dynamodb/pkg/domain/repository"

	"github.com/guregu/dynamo"
)

type UserRepository struct{}

var _ repository.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *dynamo.DB) *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Example(db *dynamo.DB) {
	//TODO
}
