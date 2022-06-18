package handler

import (
	"api/domain"
	"api/service/reply/usecase"
	"api/utils/s"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReplyHandler struct {
	Reply *usecase.ReplyUseCase
}

// add user
func (h *ReplyHandler) AddReply(c *gin.Context) {
	body := new(domain.Reply)
	if err := c.ShouldBindJSON(body); err != nil {
		s.Abort(c, err)
		return
	}
	body.Log = &domain.Log{
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := h.Reply.AddReply(body)
	if err != nil {
		s.Abort(c, err)
		return
	}
	s.Auto(c, res, err)
}

// fetch user by id
func (h *ReplyHandler) FetchReplyById(c *gin.Context) {
	_ID := c.GetHeader("ID")
	ID, err := primitive.ObjectIDFromHex(_ID)
	if err != nil {
		s.AbortWithMessage(c, "ID is not valid")
		return
	}
	res, err := h.Reply.GetReplyByID(ID)
	s.Auto(c, res, err)
}

// fetch all
func (h *ReplyHandler) FetchReply(c *gin.Context) {
	res, err := h.Reply.FetchAllReply()
	s.Auto(c, res, err)
}
