package http

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toeydevelopment/microservices-example/party-query-service/usecase"
)

func NewHTTP(usc usecase.IUsecase) error {
	g := gin.Default()

	port := os.Getenv("APP_PORT")

	portNumber := 3002

	if port == "" {
		log.Println("APP_PORT not set use default :3002")
		port = "3002"
	}

	portNumber, err := strconv.Atoi(port)

	if err != nil {
		return err
	}

	h := newHandler(usc)

	g.GET("/", h.GetAllParty)
	g.GET("/:id", h.GetPartyByID)

	return g.Run(fmt.Sprintf(":%d", portNumber))
}
