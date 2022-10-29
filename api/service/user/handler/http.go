package handler

import (
	"api/domain"
	"api/service/user/usecase"
	_hash "api/utils/password"
	"api/utils/s"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	User *usecase.UserUseCase
}

// add user
func (h *UserHandler) AddUser(c *gin.Context) {
	body := new(domain.User)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	hashPassword, err := _hash.HashPassword(body.Password)
	if err != nil {
		s.Abort(c, err)
		return
	}
	body.Password = hashPassword
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.User.AddUser(body)
	if err != nil {
		s.AbortWithMessage(c, "User already exists")
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *UserHandler) FetchUserById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.User.GetUserByID(ID)
	s.Auto(c, res, err)
}

// fetch all user
func (h *UserHandler) FetchAllUser(c *gin.Context) {
	res, err := h.User.GetAllUser()
	s.Auto(c, res, err)
}

// fetch user by username
func (h *UserHandler) FetchUserByUsername(c *gin.Context) {
	// get value username from form file
	username := c.GetHeader("username")
	if username == "" {
		s.AbortWithMessage(c, "username is required")
		return
	}
	res, err := h.User.GetUserByUsername(username)
	s.Auto(c, res, err)
}

// fetch user by email
func (h *UserHandler) FetchUserByEmail(c *gin.Context) {
	// get value email from form file
	email := c.GetHeader("email")
	if email == "" {
		s.AbortWithMessage(c, "email is required")
		return
	}
	res, err := h.User.GetUserByEmail(email)
	s.Auto(c, res, err)
}

// update by id
func (h *UserHandler) UpdateUserById(c *gin.Context) {
	body := new(domain.User)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	body.Log = &domain.Log{
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	err = h.User.UpdateUserByID(ID, body)
	if err != nil {
		s.AbortWithMessage(c, "User not found")
		return
	}
	s.AbortWithMessage(c, "Update success")
}

// update by username
func (h *UserHandler) UpdateUserByUsername(c *gin.Context) {
	body := new(domain.User)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	username := c.GetHeader("username")
	if username == "" {
		s.AbortWithMessage(c, "username is required")
		return
	}
	body.Log = &domain.Log{
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	err := h.User.UpdateUserByUsername(username, body)
	if err != nil {
		s.AbortWithMessage(c, "User not found")
		return
	}
	s.AbortWithMessage(c, "Update success")
}

// function delete
func (h *UserHandler) DeleteUserById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	err = h.User.DeleteUserByID(ID)
	if err != nil {
		s.AbortWithMessage(c, "User not found")
		return
	}
	s.AbortWithMessage(c, "Delete success")
}

// fetch pagination
func (h *UserHandler) FetchPagination(c *gin.Context) {
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
	res, err := h.User.Pagination(int64(pageInt), int64(limitInt))
	s.Auto(c, res, err)
}

// fetch search pagination
func (h *UserHandler) FetchSearchPagination(c *gin.Context) {
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
		res, err := h.User.Pagination(int64(pageInt), int64(limitInt))
		s.Auto(c, res, err)
		return
	}
	res, err := h.User.SearchPagination(int64(pageInt), int64(limitInt), search)
	s.Auto(c, res, err)
}
