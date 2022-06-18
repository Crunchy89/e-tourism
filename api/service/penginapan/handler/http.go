package handler

import (
	"api/domain"
	"api/service/penginapan/usecase"
	"api/utils/s"
	slug "api/utils/slug"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PenginapanHandler struct {
	Penginapan *usecase.PenginapanUseCase
}

// add user
func (h *PenginapanHandler) AddPenginapan(c *gin.Context) {
	body := new(domain.Penginapan)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Slug = slug.Slug(body.NamaHotel)
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Penginapan.AddPenginapan(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *PenginapanHandler) FetchPenginapanById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Penginapan.GetPenginapanByID(ID)
	s.Auto(c, res, err)
}
