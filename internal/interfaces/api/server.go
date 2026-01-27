package api

import (
	"github.com/R2Remote/ChronoGo/internal/interfaces/api/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	s := &Server{engine: r}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	jobHandler := handler.NewJobHandler()

	v1 := s.engine.Group("/api/v1")
	{
		v1.GET("/jobs", jobHandler.List)
		v1.POST("/jobs", jobHandler.Save)
		v1.DELETE("/jobs/:id", jobHandler.Delete)
		v1.POST("/jobs/:id/run", jobHandler.RunNow)
	}
}

func (s *Server) Start(addr string) error {
	return s.engine.Run(addr)
}
