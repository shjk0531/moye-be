package model

import (
	"time"
)

// Study 모델은 studies 테이블과 매핑됩니다.
type Study struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"` // study_id
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	OwnerID     uint      `gorm:"not null" json:"owner_id"` // 스터디 생성자 (User 외래키)
	Status      string    `gorm:"not null" json:"status"`   // 예: 모집 중, 진행 중, 종료 등
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// 연관관계
	Recruitment Recruitment `gorm:"foreignKey:StudyID" json:"recruitment,omitempty"`
	Users       []UserStudy `gorm:"foreignKey:StudyID" json:"users,omitempty"`
}
