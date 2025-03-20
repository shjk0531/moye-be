package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye-be/internal/domain/message/model"
	"github.com/shjk0531/moye-be/internal/domain/message/repository"
)

type Controller struct {
	repo repository.Repository
}

func NewController(repo repository.Repository) *Controller {
	return &Controller{repo: repo}
}

// RegisterRoutes는 /messages 경로 하위의 라우트를 등록합니다.
func (ctrl *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/messages")
	r.POST("/", ctrl.CreateMessage)
	r.GET("/chatroom/:chatroom_id", ctrl.GetMessagesByChatRoom)
}

func (ctrl *Controller) CreateMessage(c *gin.Context) {
	var msg model.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.repo.InsertMessage(&msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "메시지 생성 실패"})
		return
	}
	c.JSON(http.StatusCreated, msg)
}

func (ctrl *Controller) GetMessagesByChatRoom(c *gin.Context) {
	chatRoomIDStr := c.Param("chatroom_id")
	chatRoomID, err := strconv.Atoi(chatRoomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 채팅방 ID"})
		return
	}
	messages, err := ctrl.repo.FindMessagesByChatRoomID(uint(chatRoomID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "메시지 조회 실패"})
		return
	}
	c.JSON(http.StatusOK, messages)
}
