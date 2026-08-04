package network

import (
	"net/http"

	"DistributedGrep/internal/app"
	"DistributedGrep/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	service *service.GrepService
}

func NewServer(s *service.GrepService) *Server {
	return &Server{
		service: s,
	}
}

func (s *Server) Register(router *gin.Engine) {

	router.POST("/grep", s.Search)
}

func (s *Server) Search(c *gin.Context) {

	var req app.GrepRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	resp, err := s.service.Search(
		c.Request.Context(),
		req,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, resp)
}