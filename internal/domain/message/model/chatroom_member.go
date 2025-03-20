package model

import (
	"time"
)

// ChatRoomMember 모델은 채팅방의 구성원을 나타냅니다.
type ChatRoomMember struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChatRoomID uint      `gorm:"not null" json:"chatroom_id"` // ChatRoom 외래키
	UserID     uint      `gorm:"not null" json:"user_id"`     // User 외래키
	JoinedAt   time.Time `json:"joined_at"`
}
