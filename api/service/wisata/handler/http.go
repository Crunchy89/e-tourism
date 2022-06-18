package handler

import (
	"api/domain"
	"api/service/wisata/usecase"
	"api/utils/s"
	slug "api/utils/slug"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WisataHandler struct {
	Wisata *usecase.WisataUseCase
}

// add user
func (h *WisataHandler) AddWisata(c *gin.Context) {
	body := new(domain.Wisata)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Slug = slug.Slug(body.NamaObjekWisata)
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Wisata.AddWisata(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *WisataHandler) FetchWisataById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Wisata.GetWisataByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *WisataHandler) FetchWisata(c *gin.Context) {
	res, err := h.Wisata.FetchAllWisata()
	s.Auto(c, res, err)
}
