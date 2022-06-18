package handler

import (
	"api/domain"
	"api/service/visited/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VisitedHandler struct {
	Visited *usecase.VisitedUseCase
}

// add user
func (h *VisitedHandler) AddVisited(c *gin.Context) {
	body := new(domain.Visited)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Visited.AddVisited(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *VisitedHandler) FetchVisitedById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Visited.GetVisitedByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *VisitedHandler) FetchVisited(c *gin.Context) {
	res, err := h.Visited.FetchAllVisited()
	s.Auto(c, res, err)
}
