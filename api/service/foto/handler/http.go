package handler

import (
	"api/domain"
	"api/service/foto/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FotoHandler struct {
	Foto *usecase.FotoUseCase
}

// add user
func (h *FotoHandler) AddFoto(c *gin.Context) {
	body := new(domain.Foto)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Foto.AddFoto(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *FotoHandler) FetchFotoById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Foto.GetFotoByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *FotoHandler) FetchFoto(c *gin.Context) {
	res, err := h.Foto.FetchAllFoto()
	s.Auto(c, res, err)
}
