package db

import (
	"github.com/pragmatically-dev/bookstore_oauth_api/src/clients/cassandra"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"
)

type dbRepository struct{}

func NewDBRepository() IDBRepository {
	return &dbRepository{}
}

func (d *dbRepository) GetByID(ID string) (*access_token.AccessToken, *errors.APIErrors) {
	var errs errors.APIErrors
	session, err := cassandra.GetSession()
	if err != nil {
		errs.AddError(errors.NewInternalServerError("Internal Server Error", "Database Error"))
		return nil, &errs
	}
	defer session.Close()

	errs.AddError(errors.NewInternalServerError("Internal Server Error", "Database not implemented yet"))
	return nil, &errs
}
