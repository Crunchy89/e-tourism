package handler

import (
	"api/domain"
	"api/service/penginapan/usecase"
	"api/utils/s"
	slug "api/utils/slug"
	"strconv"
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

// fetch all
func (h *PenginapanHandler) FetchPenginapan(c *gin.Context) {
	res, err := h.Penginapan.FetchAllPenginapan()
	s.Auto(c, res, err)
}

// update by id
func (h *PenginapanHandler) UpdatePenginapanByID(c *gin.Context) {
	body := new(domain.Penginapan)
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
	err = h.Penginapan.UpdatePenginapanByID(ID, body)
	if err != nil {
		s.AbortWithMessage(c, "Update failed")
		return
	}
	s.AbortWithMessage(c, "Update success")
}

// delete by id
func (h *PenginapanHandler) DeletePenginapanByID(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	err = h.Penginapan.DeletePenginapanByID(ID)
	if err != nil {
		s.AbortWithMessage(c, "Delete failed")
		return
	}
	s.AbortWithMessage(c, "Delete success")
}

// fetch by slug
func (h *PenginapanHandler) FetchPenginapanBySlug(c *gin.Context) {
	slug := c.GetHeader("slug")
	if slug == "" {
		s.AbortWithMessage(c, "Slug can't be empty")
		return
	}
	res, err := h.Penginapan.GetPenginapanBySlug(slug)
	s.Auto(c, res, err)
}

// fetch pagination
func (h *PenginapanHandler) FetchPagination(c *gin.Context) {
	page := c.GetHeader("page")
	limit := c.GetHeader("limit")
	if page == "" || page == "0" {
		page = "1"
		return
	}
	if limit == "" || limit == "0" {
		limit = "10"
		return
	}
	// convert string to int using atoi
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
		return
	}
	res, err := h.Penginapan.Pagination(int64(pageInt), int64(limitInt))
	s.Auto(c, res, err)
}

// fetch search pagination
func (h *PenginapanHandler) FetchSearchPagination(c *gin.Context) {
	page := c.GetHeader("page")
	limit := c.GetHeader("limit")
	search := c.GetHeader("search")
	if page == "" || page == "0" {
		page = "1"
		return
	}
	if limit == "" || limit == "0" {
		limit = "10"
		return
	}

	// convert string to int using atoi
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
		return
	}
	if search == "" {
		res, err := h.Penginapan.Pagination(int64(pageInt), int64(limitInt))
		s.Auto(c, res, err)
		return
	}
	res, err := h.Penginapan.SearchPagination(int64(pageInt), int64(limitInt), search)
	s.Auto(c, res, err)
}
