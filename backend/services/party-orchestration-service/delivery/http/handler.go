package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/constant"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/dto"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/usecase"
)

type handler struct {
	usc usecase.IUsecase
}

func newHandler(usc usecase.IUsecase) handler {
	return handler{usc}
}

func (h handler) CreateParty(c *gin.Context) {
	var body CreatePartyRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	partyID, err := h.usc.CreateParty(c.Request.Context(), dto.CreatePartyDTO{
		Name:        body.Name,
		Description: body.Description,
		SeatLimit:   body.SeatLimit,
		UserEmail:   c.GetString(constant.UserEmail),
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"erorr": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"party_id": partyID,
		},
	})

}

func (h handler) UpdateParty(c *gin.Context) {}

func (h handler) DeleteParty(c *gin.Context) {}

func (h handler) JoinParty(c *gin.Context) {}

func (h handler) UnjoinParty(c *gin.Context) {}
