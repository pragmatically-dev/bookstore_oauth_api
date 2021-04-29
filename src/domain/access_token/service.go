package access_token

import (
	"strings"

	"github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"
)

type Service interface {
	GetByID(string) (*AccessToken, *errors.APIErrors)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.APIErrors) {
	var errs errors.APIErrors
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		errs.AddError(errors.NewInternalServerError("Invalid access token", "Internal Server Error"))
		return nil, &errs
	}
	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
