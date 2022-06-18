package handler

import (
	"api/domain"
	"api/service/travel/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TravelHandler struct {
	Travel *usecase.TravelUseCase
}

// add user
func (h *TravelHandler) AddTravel(c *gin.Context) {
	body := new(domain.Travel)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Travel.AddTravel(body)
	s.Auto(c, res, err)
}

// fetch all
func (h *TravelHandler) FetchTravel(c *gin.Context) {
	res, err := h.Travel.FetchAllTravel()
	s.Auto(c, res, err)
}

// fetch user by id
func (h *TravelHandler) FetchTravelById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Travel.GetTravelByID(ID)
	s.Auto(c, res, err)
}
