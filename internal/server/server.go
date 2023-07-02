package server

import (
	"time"

	"github.com/DogGoOrg/doggo-api-gateway/internal/endpoints"
	"github.com/DogGoOrg/doggo-api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"golang.org/x/exp/slog"
)

type Server struct {
	engine *gin.Engine
	logger *slog.Logger
}

func NewServer(logger *slog.Logger) *Server {
	return &Server{
		engine: gin.New(),
		logger: logger,
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
	server.engine.Use(
		gin.Recovery(),
	)
	return server.engine.Run(":" + addr)
}

func (server *Server) Engine() *gin.Engine {
	return server.engine
}

func ConfigureRoutes(server *Server) {

	loggerMiddleware := middleware.NewLoggerMiddleware(server.logger)

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

	eng.Use(loggerMiddleware)

	eng.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "check",
		})
	})

	eng.GET("/ping_account", endpoints.PingAccountHandler)
	eng.GET("/ping_tracker", endpoints.TrackerPingHandler)
}
