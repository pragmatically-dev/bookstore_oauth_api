package app

import (
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/repository/db"
)

//InitApplication as
func InitApplication() {
	dbRepository := db.DBRepository
	atService := access_token.NewService()
}
