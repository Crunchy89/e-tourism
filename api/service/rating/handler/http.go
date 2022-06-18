package handler

import (
	"api/domain"
	"api/service/rating/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingHandler struct {
	Rating *usecase.RatingUseCase
}

// add user
func (h *RatingHandler) AddRating(c *gin.Context) {
	body := new(domain.Rating)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Rating.AddRating(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *RatingHandler) FetchRatingById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Rating.GetRatingByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *RatingHandler) FetchRating(c *gin.Context) {
	res, err := h.Rating.FetchAllRating()
	s.Auto(c, res, err)
}
