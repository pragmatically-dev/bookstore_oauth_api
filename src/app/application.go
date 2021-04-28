package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/http/http_access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/repository/db"
)

var (
	router = gin.Default()
)

//InitApplication as
func InitApplication() {

	atService := access_token.NewService(db.NewDBRepository())
	atHandler := http_access_token.NewHandler(atService)
	initializeRoutes(atHandler)
	router.Run()
}
