package handler

import (
	"api/domain"
	"api/service/kamar/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KamarHandler struct {
	Kamar *usecase.KamarUseCase
}

// add user
func (h *KamarHandler) AddKamar(c *gin.Context) {
	body := new(domain.Kamar)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Kamar.AddKamar(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *KamarHandler) FetchKamarById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Kamar.GetKamarByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *KamarHandler) FetchKamar(c *gin.Context) {
	res, err := h.Kamar.FetchAllKamar()
	s.Auto(c, res, err)
}
