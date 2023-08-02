package endpoints

import (
	"github.com/DogGoOrg/doggo-api-gateway/internal/helpers"
	"github.com/gin-gonic/gin"
)

func Refresh(ctx *gin.Context) {
	// var email, password string
	conn, err := GrpcController.ConnGrpc("ACCOUNT_HOST")

	if err != nil {
		helpers.Error5xx(ctx, err)
		return
	}

	defer conn.Close()

	// client := Account.NewAccountClient(conn)

	// res, err := client.Login(context.TODO(), &Account.LoginRequest{})
}
