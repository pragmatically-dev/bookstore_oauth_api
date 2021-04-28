package app

import "github.com/pragmatically-dev/bookstore_oauth_api/src/http/http_access_token"

func initializeRoutes(handler http_access_token.AccessTokenHandler) {
	router.GET("/oauth/access_token/:access_token_id", handler.GetByID)
}
