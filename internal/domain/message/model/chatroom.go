package model

import (
	"time"
)

// ChatRoom 모델은 채팅방 정보를 저장합니다.
type ChatRoom struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"` // chatroom_id
	Type      string    `gorm:"not null" json:"type"`               // 채팅방 종류 (private, group, video)
	StudyID   *uint     `json:"study_id"`                           // 스터디와 연계된 경우 (nullable)
	CreatedAt time.Time `json:"created_at"`

	// 연관관계
	Members      []ChatRoomMember `gorm:"foreignKey:ChatRoomID" json:"members,omitempty"`
	Messages     []Message        `gorm:"foreignKey:ChatRoomID" json:"messages,omitempty"`
	VideoSession []VideoChatSession `gorm:"foreignKey:ChatRoomID" json:"video_sessions,omitempty"`
}
