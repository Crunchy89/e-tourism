package handler

import (
	"api/domain"
	"api/service/user/usecase"
	"api/utils/s"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	User *usecase.UserUseCase
}

// add submission
func (h *UserHandler) AddUser(c *gin.Context) {
	body := new(domain.User)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	res, err := h.User.AddUser(body)
	s.Auto(c, res, err)
}

// fetch submissions by event id
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
	res, err := h.User.GetUserByUsername(username)
	s.Auto(c, res, err)
}

// fetch user by token
func (h *UserHandler) FetchUserByToken(c *gin.Context) {
	// get value token from form file
	token := c.GetHeader("token")
	res, err := h.User.GetUserByToken(token)
	s.Auto(c, res, err)
}
