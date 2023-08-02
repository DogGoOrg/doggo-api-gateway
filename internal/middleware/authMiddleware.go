package middleware

import (
	"strings"

	"github.com/DogGoOrg/doggo-api-gateway/internal/dto"
	"github.com/DogGoOrg/doggo-api-gateway/internal/endpoints"
	"github.com/DogGoOrg/doggo-api-gateway/internal/helpers"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Account"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		grpcController := endpoints.GrpcController

		if authHeader == "" {
			helpers.UnauthorizedError(ctx)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == authHeader {
			helpers.UnauthorizedError(ctx)
			return
		}

		conn, err := grpcController.ConnGrpc("ACCOUNT_HOST")

		if err != nil {
			helpers.Error5xx(ctx, err)
			return
		}

		defer conn.Close()

		client := Account.NewAccountClient(conn)

		res, err := client.CheckAuthorization(ctx, &Account.CheckAuthorizationRequest{AccessToken: token})

		if err != nil {
			errStatus, _ := status.FromError(err)
			helpers.Error5xx(ctx, errStatus.Err())
			return
		}

		dto := dto.TokenDTO{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		}

		ctx.Next()
	}
}
