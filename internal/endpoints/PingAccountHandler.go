package endpoints

import (
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Account"
	"github.com/gin-gonic/gin"
)

func PingAccountHandler(ctx *gin.Context, res *Account.PingResponse) {
	ctx.JSON(200, res)
}
