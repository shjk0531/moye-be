package main

import (
	"fmt"
	"log"

	"time"

	messageModel "github.com/shjk0531/moye-be/internal/domain/message/model"
	notificationModel "github.com/shjk0531/moye-be/internal/domain/notification/model"
	studyModel "github.com/shjk0531/moye-be/internal/domain/study/model"
	userModel "github.com/shjk0531/moye-be/internal/domain/user/model"

	"github.com/shjk0531/moye-be/internal/global/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 설정 로드
	config.InitConfig()

	// DSN 생성 (SSL 모드는 개발환경에 맞게 disable)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		config.Config.DBHost,
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBName,
		config.Config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}

	// 각 도메인 모델의 마이그레이션 실행
	err = db.AutoMigrate(
		&userModel.User{},
		&studyModel.Study{},
		&studyModel.Recruitment{},
		&studyModel.UserStudy{},
		&messageModel.ChatRoom{},
		&messageModel.ChatRoomMember{},
		&messageModel.Message{},
		&messageModel.VideoChatSession{},
		&notificationModel.Notification{},
	)
	if err != nil {
		log.Fatalf("마이그레이션 실패: %v", err)
	}

	log.Println("데이터베이스 마이그레이션 성공!")

	// (선택) 예시 데이터 삽입: 간단한 시간 값 테스트
	log.Println("현재 시간:", time.Now())
}
