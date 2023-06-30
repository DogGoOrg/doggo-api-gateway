package endpoints

import (
	"context"

	"github.com/DogGoOrg/doggo-api-gateway/internal/dto"
	"github.com/DogGoOrg/doggo-api-gateway/internal/utils"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Tracker"
	"github.com/gin-gonic/gin"
)

func TrackerPingHandler(ctx *gin.Context) {
	conn, err := grpcController.ConnGrpc("TRACKER_HOST")

	if err != nil {
		utils.Error5xx(ctx, err)
		return
	}

	defer conn.Close()

	client := Tracker.NewTrackerClient(conn)

	res, err := client.Ping(context.Background(), &Tracker.PingRequest{})

	if err != nil {
		utils.Error5xx(ctx, err)
		return
	}

	dto := &dto.TrackerPingDTO{Status: res.Status}

	response := utils.ResponseWrapper{Status: true, Error: nil, Data: dto}

	ctx.JSON(200, response)
}
