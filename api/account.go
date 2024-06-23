package api

import (
	"net/http"

	db "github.com/christy712/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateAccountParams struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD INR"`
}

func (Server *Server) CreateAccount(ctx *gin.Context) {

	var req CreateAccountParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return

	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}

	account, err := Server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
