package server

import (
	"context"
	"log"
	"time"

	"github.com/DogGoOrg/doggo-api-gateway/internal/endpoints"
	"github.com/DogGoOrg/doggo-api-gateway/proto/proto_services/Account"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}

func (server *Server) Run(addr string) error {
	server.engine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          500 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	return server.engine.Run(":" + addr)
}

func (server *Server) Engine() *gin.Engine {
	return server.engine
}

func (server *Server) DialAccountServiceConnection(addr string) (Account.AccountClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		conn.Close()
		return nil, err
	}

	client := Account.NewAccountClient(conn)
	return client, nil
}

func ConfigureRoutes(server *Server, conn *ServiceConnections) {

	eng := server.engine

	eng.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
			return
		}
		ctx.Next()
	})

	eng.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "check",
		})
	})

	eng.GET("/ping_account", func(ginCtx *gin.Context) {
		ctx := context.Background()
		ctx, cancelFn := context.WithTimeout(ctx, time.Second*15)
		defer cancelFn()

		service := *conn.AccountService

		res, err := service.Ping(ctx, &Account.PingRequest{})

		if err != nil {
			log.Fatalln(err)
		}

		endpoints.PingAccountHandler(ginCtx, res)
	})
}
