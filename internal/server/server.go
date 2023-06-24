package server

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
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
