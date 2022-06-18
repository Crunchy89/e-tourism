package handler

import (
	"api/domain"
	"api/service/searched/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchedHandler struct {
	Searched *usecase.SearchedUseCase
}

// add user
func (h *SearchedHandler) AddSearched(c *gin.Context) {
	body := new(domain.Searched)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Searched.AddSearched(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *SearchedHandler) FetchSearchedById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Searched.GetSearchedByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *SearchedHandler) FetchSearched(c *gin.Context) {
	res, err := h.Searched.FetchAllSearched()
	s.Auto(c, res, err)
}
