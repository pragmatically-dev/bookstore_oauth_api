package http_access_token

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/domain/access_token"
	"github.com/pragmatically-dev/bookstore_oauth_api/src/utils/errors"
)

//AccessTokenHandler Interface that defines the methods of a handler
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
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
	accessTokenID := ctx.Param("access_token_id")
	accessToken, err := h.service.GetByID(accessTokenID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)

}

func (h *accessTokenHandler) Create(ctx *gin.Context) {
	var at access_token.AccessToken
	err := ctx.BindJSON(&at)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.NewBadRequestError("Bad Request", err.Error()))
		return
	}

	errs := h.service.Create(&at)
	if errs != nil {
		fmt.Println(&errs)
		ctx.JSON(http.StatusInternalServerError, errs)
		return
	}
	ctx.JSON(http.StatusOK, at)
}
