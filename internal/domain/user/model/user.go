package model

import (
	"time"
)

// User 모델은 users 테이블과 매핑됩니다.
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`          // user_id
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`                            // 보안을 위해 json에 노출하지 않음
	ProfileInfo  string    `json:"profile_info"`                                 // 프로필(사진 URL, 소개 등)
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// 연관관계 (옵션)
	// Studies: 사용자가 생성한 스터디 (1:N)
	// Notifications: 사용자에게 전달된 알림 (1:N)
}
