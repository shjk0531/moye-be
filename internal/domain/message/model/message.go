package model

import "time"

// Message 모델은 채팅 메시지를 저장합니다.
type Message struct {
	// MongoDB는 기본적으로 ObjectID를 사용하지만, 단순화를 위해 string으로 처리합니다.
	ID          string    `json:"id" bson:"_id,omitempty"`
	ChatRoomID  uint      `json:"chatroom_id" bson:"chatroom_id"`
	SenderID    uint      `json:"sender_id" bson:"sender_id"`
	Content     string    `json:"content" bson:"content"`
	MessageType string    `json:"message_type" bson:"message_type"`
	SentAt      time.Time `json:"sent_at" bson:"sent_at"`
}
