package endpoints

import (
	"errors"

	"github.com/DogGoOrg/doggo-api-gateway/internal/dto"
	"github.com/DogGoOrg/doggo-api-gateway/internal/helpers"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Account"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type registerReqBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context) {
	registerReqBody := &registerReqBody{}

	//bind body to struct
	if err := ctx.BindJSON(registerReqBody); err != nil {
		helpers.Error5xx(ctx, err)
		return
	}

	email, password := registerReqBody.Email, registerReqBody.Password

	if email == "" || password == "" {
		helpers.Error5xx(ctx, errors.New("invalid request body"))
		return
	}

	conn, err := grpcController.ConnGrpc("ACCOUNT_HOST")

	if err != nil {
		helpers.Error5xx(ctx, err)
		return
	}

	defer conn.Close()

	client := Account.NewAccountClient(conn)

	res, err := client.Register(ctx, &Account.RegisterRequest{Email: email, Password: password})

	if err != nil {
		errStatus, _ := status.FromError(err)
		helpers.Error5xx(ctx, errStatus.Err())
		return
	}

	dto := &dto.RegisterDTO{Status: res.Status}

	//TODO: send activation email

	response := helpers.ResponseWrapper{Success: true, Error: nil, Data: dto}

	ctx.JSON(201, response)
}
