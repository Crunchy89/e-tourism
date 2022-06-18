package handler

import (
	"api/domain"
	"api/service/review/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewHandler struct {
	Review *usecase.ReviewUseCase
}

// add user
func (h *ReviewHandler) AddReview(c *gin.Context) {
	body := new(domain.Review)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Review.AddReview(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *ReviewHandler) FetchReviewById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Review.GetReviewByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *ReviewHandler) FetchReview(c *gin.Context) {
	res, err := h.Review.FetchAllReview()
	s.Auto(c, res, err)
}
