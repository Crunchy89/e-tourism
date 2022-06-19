package handler

import (
	"api/domain"
	"api/service/viewed/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ViewedHandler struct {
	Viewed *usecase.ViewedUseCase
}

// add user
func (h *ViewedHandler) AddViewed(c *gin.Context) {
	body := new(domain.Viewed)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Viewed.AddViewed(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *ViewedHandler) FetchViewedById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Viewed.GetViewedByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *ViewedHandler) FetchViewed(c *gin.Context) {
	res, err := h.Viewed.FetchAllViewed()
	s.Auto(c, res, err)
}
