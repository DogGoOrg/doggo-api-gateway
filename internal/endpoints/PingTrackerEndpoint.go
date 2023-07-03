package endpoints

import (
	"context"

	"github.com/DogGoOrg/doggo-api-gateway/internal/dto"
	"github.com/DogGoOrg/doggo-api-gateway/internal/helpers"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Tracker"
	"github.com/gin-gonic/gin"
)

func TrackerPingHandler(ctx *gin.Context) {
	conn, err := grpcController.ConnGrpc("TRACKER_HOST")

	if err != nil {
		helpers.Error5xx(ctx, err)
		return
	}

	defer conn.Close()

	client := Tracker.NewTrackerClient(conn)

	res, err := client.Ping(context.Context(ctx), &Tracker.PingRequest{})

	if err != nil {
		helpers.Error5xx(ctx, err)
		return
	}

	dto := &dto.TrackerPingDTO{Status: res.Status}

	response := helpers.ResponseWrapper{Success: true, Error: nil, Data: dto}

	ctx.JSON(200, response)
}
