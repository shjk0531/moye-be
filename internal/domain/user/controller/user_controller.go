package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye-be/internal/domain/user/model"
	"github.com/shjk0531/moye-be/internal/domain/user/service"
)

type Controller struct {
	service service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{service: s}
}

// RegisterRoutes는 /users 경로 하위의 라우트를 등록합니다.
func (ctrl *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/users")
	r.POST("/", ctrl.CreateUser)
	r.GET("/:id", ctrl.GetUser)
}

func (ctrl *Controller) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "사용자 생성 실패"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (ctrl *Controller) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID"})
		return
	}

	user, err := ctrl.service.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "사용자를 찾을 수 없음"})
		return
	}

	c.JSON(http.StatusOK, user)
}
