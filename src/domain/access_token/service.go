package access_token

import "github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"

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

func (s *service) GetByID(string) (*AccessToken, *errors.APIErrors) {
	return nil, nil
}
