package endpoints

import (
	"context"

	"github.com/DogGoOrg/doggo-api-gateway/internal/dto"
	"github.com/DogGoOrg/doggo-api-gateway/internal/helpers"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Account"
	"github.com/gin-gonic/gin"
)

func PingAccountHandler(ctx *gin.Context) {
	conn, err := grpcController.ConnGrpc("ACCOUNT_HOST")

	if err != nil {
		helpers.Error5xx(ctx, err)
		return
	}

	defer conn.Close()

	client := Account.NewAccountClient(conn)

	res, err := client.Ping(context.Context(ctx), &Account.PingRequest{})

	if err != nil {
		helpers.Error5xx(ctx, err)
		return
	}

	dto := &dto.AccountPingDTO{Status: res.Status}

	response := helpers.ResponseWrapper{Success: true, Error: nil, Data: dto}

	ctx.JSON(200, response)
}
