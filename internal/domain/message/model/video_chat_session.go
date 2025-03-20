package model

import (
	"time"
)

// VideoChatSession 모델은 화상채팅 세션 정보를 저장합니다.
type VideoChatSession struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"` // session_id
	ChatRoomID uint      `gorm:"not null" json:"chatroom_id"`        // 채팅방 외래키 (화상채팅 전용 또는 그룹채팅 연계)
	StartedAt  time.Time `json:"started_at"`
	EndedAt    time.Time `json:"ended_at"`
	SessionURL string    `json:"session_url"`  // 접속 URL 또는 세션 정보
	CreatedBy  uint      `gorm:"not null" json:"created_by"` // 세션을 시작한 사용자 (User 외래키)
}
