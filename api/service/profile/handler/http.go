package handler

import (
	"api/domain"
	"api/service/profile/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileHandler struct {
	Profile *usecase.ProfileUseCase
}

// add user
func (h *ProfileHandler) AddProfile(c *gin.Context) {
	body := new(domain.Profile)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Profile.AddProfile(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *ProfileHandler) FetchProfileById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Profile.GetProfileByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *ProfileHandler) FetchProfile(c *gin.Context) {
	res, err := h.Profile.FetchAllProfile()
	s.Auto(c, res, err)
}
