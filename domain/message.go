package domain

import "time"

type Message struct {
	ID         string    `json:"id"`
	SenderID   string    `json:"senderId"`
	ReceiverID string    `json:"ReceiverId"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"createdAt"`
}
