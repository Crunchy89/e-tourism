package handler

import (
	"api/domain"
	"api/service/fasilitas/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FasilitasHandler struct {
	Fasilitas *usecase.FasilitasUseCase
}

// add user
func (h *FasilitasHandler) AddFasilitas(c *gin.Context) {
	body := new(domain.Fasilitas)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Fasilitas.AddFasilitas(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *FasilitasHandler) FetchFasilitasById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Fasilitas.GetFasilitasByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *FasilitasHandler) FetchFasilitas(c *gin.Context) {
	res, err := h.Fasilitas.FetchAllFasilitas()
	s.Auto(c, res, err)
}
