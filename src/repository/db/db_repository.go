package db

import (
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"
)

type dbRepository struct {
}

func NewDBRepository() DBRepository {
	return &dbRepository{}
}

func (d *dbRepository) GetByID(string) (*access_token.AccessToken, *errors.APIErrors) {
	return nil, nil
}
