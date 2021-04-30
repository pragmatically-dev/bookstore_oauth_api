package db

import (
	"github.com/gocql/gocql"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/clients/cassandra"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"
)

const (
	queryGetAccessToken = `SELECT access_token,user_id,client_id,expires
	 FROM access_tokens
	 WHERE access_token = ?`

	queryCreateAccessToken = `INSERT INTO access_tokens(access_token,user_id,client_id,expires)
		VALUES(?,?,?,?);
	`
	queryUpdateAccessToken = `UPDATE access_tokens SET expires = ? WHERE access_token =?;
	`
)

type dbRepository struct{}

func NewDBRepository() IDBRepository {
	return &dbRepository{}
}

func (d *dbRepository) GetByID(ID string) (*access_token.AccessToken, *errors.APIErrors) {
	var errs errors.APIErrors
	var result access_token.AccessToken
	err := cassandra.GetSession().Query(queryGetAccessToken, ID).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires)
	if err != nil {
		if err == gocql.ErrNotFound {
			errs.AddError(errors.NewNotFoundError("Not Found", "No Access Token found with given ID"))
			return nil, &errs
		}
		errs.AddError(errors.NewInternalServerError("Internal Server Error", err.Error()))
		return nil, &errs
	}

	return &result, nil
}

func (d *dbRepository) Create(accessToken *access_token.AccessToken) *errors.APIErrors {
	var errs errors.APIErrors
	err := cassandra.GetSession().Query(queryCreateAccessToken,
		accessToken.AccessToken,
		accessToken.UserID,
		accessToken.ClientID,
		accessToken.Expires,
	).Exec()
	if err != nil {
		errs.AddError(errors.NewInternalServerError("Internal Server Error", err.Error()))
		return &errs
	}

	return nil
}

func (d *dbRepository) UpdateExpirationTime(accessToken *access_token.AccessToken) *errors.APIErrors {
	var errs errors.APIErrors
	err := cassandra.GetSession().Query(queryUpdateAccessToken,
		accessToken.Expires,
		accessToken.AccessToken,
	).Exec()
	if err != nil {
		errs.AddError(errors.NewInternalServerError("Internal Server Error", err.Error()))
		return &errs
	}

	return nil
}
