package http

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/usecase"
)

func NewHTTP(usc usecase.IUsecase, authMid gin.HandlerFunc) error {

	g := gin.Default()

	port := os.Getenv("APP_PORT")

	portNumber := 8080

	if port == "" {
		log.Println("APP_PORT not set use default :8080")
		port = "8080"
	}

	portNumber, err := strconv.Atoi(port)

	if err != nil {
		return err
	}

	h := newHandler(usc)

	g.POST("/", h.CreateParty)
	g.PATCH("/:id", authMid, h.UpdateParty)
	g.DELETE("/:id", authMid, h.DeleteParty)
	g.POST("/:id/join", authMid, h.JoinParty)
	g.POST("/:id/unjoin", authMid, h.UnjoinParty)

	return g.Run(fmt.Sprintf(":%d", portNumber))
}
