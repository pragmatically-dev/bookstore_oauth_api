package db

import (
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.APIErrors)
}
