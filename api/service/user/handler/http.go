package handler

import (
	"api/domain"
	"api/service/user/usecase"
	_hash "api/utils/password"
	"api/utils/s"

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
	res, err := h.User.AddUser(body)
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
	username := c.Request.FormValue("username")
	if username == "" {
		s.AbortWithMessage(c, "username is required")
		return
	}
	res, err := h.User.GetUserByUsername(username)
	s.Auto(c, res, err)
}

// fetch user by token
func (h *UserHandler) FetchUserByToken(c *gin.Context) {
	// get value token from form file
	token := c.GetHeader("token")
	if token == "" {
		s.AbortWithMessage(c, "token is required")
		return
	}
	res, err := h.User.GetUserByToken(token)
	s.Auto(c, res, err)
}
