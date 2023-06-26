package endpoints

import (
	"context"
	"log"

	"github.com/DogGoOrg/doggo-api-gateway/internal/utils"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Account"
	"github.com/gin-gonic/gin"
)

var (
	handler = new(utils.GrpcConnectionController)
)

func PingAccountHandler(ctx *gin.Context) {
	conn, err := handler.ConnGrpc("ACCOUNT_HOST")

	if err != nil {
		log.Println(err)
	}

	client := Account.NewAccountClient(conn)

	res, err := client.Ping(context.Background(), &Account.PingRequest{})

	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"status": false, "error": err})
		return
	}

	ctx.JSON(200, res)

}
