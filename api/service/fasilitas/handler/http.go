package handler

import (
	"api/domain"
	"api/service/fasilitas/usecase"
	"api/utils/s"
	"strconv"
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

// delete by id
func (h *FasilitasHandler) DeleteFasilitas(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	err = h.Fasilitas.DeleteFasilitasByID(ID)
	if err != nil {
		s.AbortWithMessage(c, "Failed to delete")
		return
	}
	s.AbortWithMessage(c, "Successfully deleted")
}

// update by id
func (h *FasilitasHandler) UpdateFasilitas(c *gin.Context) {
	body := new(domain.Fasilitas)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	err = h.Fasilitas.UpdateFasilitasByID(ID, body)
	if err != nil {
		s.AbortWithMessage(c, "Failed to update")
		return
	}
	s.AbortWithMessage(c, "Successfully updated")
}

// fetch all
func (h *FasilitasHandler) FetchFasilitas(c *gin.Context) {
	res, err := h.Fasilitas.FetchAllFasilitas()
	s.Auto(c, res, err)
}

// fetch by foreign id
func (h *FasilitasHandler) FetchFasilitasByForeignID(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Fasilitas.GetFasilitasByForeignID(ID)
	s.Auto(c, res, err)
}

// fetch search pagination
func (h *FasilitasHandler) FetchFasilitasSearchPagination(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	if page == "" || page == "0" {
		page = "1"
	}
	if limit == "" || limit == "0" {
		limit = "10"
	}
	// convert to int using atoi
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	// get foreignID
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	search := c.Query("search")
	if search == "" {
		res, err := h.Fasilitas.GetFasilitasByForeignIDPagination(ID, int64(pageInt), int64(limitInt))
		s.Auto(c, res, err)
		return
	}
	res, err := h.Fasilitas.GetFasilitasSearchByForeignIDPagination(ID, int64(pageInt), int64(limitInt), search)
	s.Auto(c, res, err)
}
