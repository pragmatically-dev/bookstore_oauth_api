package http_access_token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
)

//AccessTokenHandler Interface that defines the methods of a handler
type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "Not Implemented")
}
