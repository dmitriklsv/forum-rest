package entity

import "time"

type Session struct {
	ID         uint64    `json:"id,omitempty"`
	UserID     uint64    `json:"user_id,omitempty"`
	Token      string    `json:"token,omitempty"`
	ExpireTime time.Time `json:"expire_time,omitempty"`
}
