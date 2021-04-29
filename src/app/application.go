package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/clients/cassandra"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/http/http_access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/repository/db"
)

var (
	router = gin.Default()
)

//InitApplication as
func InitApplication() {
	//DB initialize
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	//Registering services
	atService := access_token.NewService(db.NewDBRepository())
	atHandler := http_access_token.NewHandler(atService)
	//setting up the middlewares
	router.Use(Json())
	//Initalize the routes
	initializeRoutes(atHandler)

	router.Run()
}

//Middleware Json
func Json() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
	}
}
