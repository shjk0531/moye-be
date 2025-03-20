package routes

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	// PostgreSQL: User 도메인
	userContainer "github.com/shjk0531/moye-be/internal/domain/user/controller"
	userRepository "github.com/shjk0531/moye-be/internal/domain/user/repository"
	userService "github.com/shjk0531/moye-be/internal/domain/user/service"

	// MongoDB: Message 도메인
	messageContainer "github.com/shjk0531/moye-be/internal/domain/message/controller"
	messageRepository "github.com/shjk0531/moye-be/internal/domain/message/repository"

	"github.com/shjk0531/moye-be/internal/global/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine) {
	// PostgreSQL 연결 (정형 데이터: User, Study, Notification 등)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		config.Config.DBHost,
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBName,
		config.Config.DBPort,
	)
	pgDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("PostgreSQL 연결 실패: " + err.Error())
	}

	api := router.Group("/api/v1")

	// User 도메인 라우트 등록 (PostgreSQL 사용)
	userRepo := userRepository.NewRepository(pgDB)
	userServ := userService.NewService(userRepo)
	userCtrl := userContainer.NewController(userServ)
	userCtrl.RegisterRoutes(api)

	// MongoDB 연결 (메시지 데이터: 채팅, 로그 등)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Config.MongoURI))
	if err != nil {
		panic("MongoDB 연결 실패: " + err.Error())
	}
	// 연결 확인
	if err := mongoClient.Ping(ctx, nil); err != nil {
		panic("MongoDB ping 실패: " + err.Error())
	}

	// Message 도메인 라우트 등록 (MongoDB 사용)
	msgRepo := messageRepository.NewRepository(mongoClient, config.Config.MongoDB)
	msgCtrl := messageContainer.NewController(msgRepo)
	msgCtrl.RegisterRoutes(api)
}
