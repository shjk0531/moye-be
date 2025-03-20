package model

import (
	"time"
)

// Notification 모델은 알림 정보를 저장합니다.
type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"` // alert_id
	UserID    uint      `gorm:"not null" json:"user_id"`            // 알림 수신자 (User 외래키)
	Message   string    `json:"message"`                            // 알림 내용
	AlertType string    `json:"alert_type"`                         // 예: 채팅, 모집, 시스템 등
	Status    string    `json:"status"`                             // 확인 상태 (읽음, 미확인 등)
	CreatedAt time.Time `json:"created_at"`
}
