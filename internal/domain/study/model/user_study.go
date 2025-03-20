package model

import (
	"time"
)

// UserStudy는 사용자와 스터디 간 다대다 관계를 해결하는 조인 테이블입니다.
type UserStudy struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudyID  uint      `gorm:"not null" json:"study_id"` // 스터디 외래키
	UserID   uint      `gorm:"not null" json:"user_id"`  // 사용자 외래키
	Role     string    `json:"role"`                   // 예: 관리자, 참여자
	JoinedAt time.Time `json:"joined_at"`
}
