package access_token

import (
	"time"

	"gopkg.in/validator.v2"

	"github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token" validate:"nonzero"`
	UserID      int64  `json:"user_id" "validate:"nonzero"`
	ClientID    int64  `json:"client_id" validate:"nonzero"`
	Expires     int64  `json:"expires" validate:"nonzero"`
}

func (at *AccessToken) Validate() *errors.APIErrors {
	var errs errors.APIErrors
	err := validator.Validate(at)
	if err != nil {
		errs.AddError(errors.NewBadRequestError("Bad Request", err.Error()))
		return &errs
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
