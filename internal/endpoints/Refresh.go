package endpoints

import (
	"github.com/DogGoOrg/doggo-api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
)

func Refresh(ctx *gin.Context) {
	// var email, password string
	conn, err := grpcController.ConnGrpc("ACCOUNT_HOST")

	if err != nil {
		utils.Error5xx(ctx, err)
		return
	}

	defer conn.Close()

	// client := Account.NewAccountClient(conn)

	// res, err := client.Login(context.TODO(), &Account.LoginRequest{})
}
