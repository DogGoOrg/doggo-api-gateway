package endpoints

import (
	"context"
	"log"

	"github.com/DogGoOrg/doggo-api-gateway/internal/dto"
	"github.com/DogGoOrg/doggo-api-gateway/internal/utils"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Tracker"
	"github.com/gin-gonic/gin"
)

func TrackerPingHandler(ctx *gin.Context) {
	conn, err := handler.ConnGrpc("TRACKER_HOST")

	if err != nil {
		log.Println(err)
	}

	client := Tracker.NewTrackerClient(conn)

	res, err := client.Ping(context.Background(), &Tracker.PingRequest{})

	if err != nil {
		response := utils.ResponseWrapper{
			Status: false,
			Error:  err,
			Data:   nil,
		}

		ctx.AbortWithStatusJSON(500, response)
		return
	}

	dto := dto.TrackerPingDTO{Status: res.Status}

	response := utils.ResponseWrapper{true, nil, dto}

	ctx.JSON(200, response)
}
