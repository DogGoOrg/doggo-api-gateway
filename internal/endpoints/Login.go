package endpoints

import (
	"context"
	"errors"

	"github.com/DogGoOrg/doggo-api-gateway/internal/dto"
	"github.com/DogGoOrg/doggo-api-gateway/internal/utils"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Account"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type loginReqBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var reqBody loginReqBody

	if err := ctx.BindJSON(&reqBody); err != nil {
		utils.Error5xx(ctx, err)
		return
	}

	email, password := reqBody.Email, reqBody.Password

	if email != "" || password != "" {
		utils.Error5xx(ctx, errors.New("invalid request body"))
		return
	}

	conn, err := grpcController.ConnGrpc("ACCOUNT_HOST")

	if err != nil {
		utils.Error5xx(ctx, err)
		return
	}

	defer conn.Close()

	defer conn.Close()

	client := Account.NewAccountClient(conn)

	res, err := client.Login(context.TODO(), &Account.LoginRequest{Email: email, Password: password})

	if err != nil {
		errStatus, _ := status.FromError(err)
		utils.Error5xx(ctx, errStatus.Err())
		return
	}

	dto := &dto.LoginDto{
		Id:           res.Id,
		Email:        res.Email,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}

	response := utils.ResponseWrapper{Status: true, Error: nil, Data: dto}

	ctx.JSON(200, response)
}
