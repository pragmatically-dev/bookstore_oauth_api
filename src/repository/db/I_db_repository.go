package db

import (
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"
)

type IDBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.APIErrors)
	Create(*access_token.AccessToken) *errors.APIErrors
	UpdateExpirationTime(*access_token.AccessToken) *errors.APIErrors
}
