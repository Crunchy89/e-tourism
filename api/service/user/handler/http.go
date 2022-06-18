package handler

import (
	"api/domain"
	"api/service/user/usecase"
	_hash "api/utils/password"
	"api/utils/s"
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
