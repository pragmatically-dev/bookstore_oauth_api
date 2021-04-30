package access_token

import (
	"github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.APIErrors)
	Create(*AccessToken) *errors.APIErrors
	UpdateExpirationTime(*AccessToken) *errors.APIErrors
}
