package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toeydevelopment/microservices-example/party-query-service/usecase"
)

type handler struct {
	usc usecase.IUsecase
}

func newHandler(usc usecase.IUsecase) handler {
	return handler{usc}
}

func (h handler) GetAllParty(c *gin.Context) {

	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 10
	}

	party, cursor, err := h.usc.FindParties(c.Request.Context(), c.Query("cursor"), int64(limit))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   party,
		"cursor": cursor,
	})
}

func (h handler) GetPartyByID(c *gin.Context) {

	id := c.Param("id")

	fmt.Println(id)

	party, err := h.usc.FindPartiyByID(c.Request.Context(), id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": party,
	})
}
