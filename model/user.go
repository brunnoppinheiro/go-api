package model

import "time"

type User struct {
	UserID      int       `json:"user_id"`
	UserName    string    `json:"user_name"`
	UserAge     int       `json:"user_age"`
	UserEmail   string    `json:"user_email"`
	UserCreated time.Time `json:"user_created"`
	UserPhone   string    `json:"user_phone"`
	UserStatus  bool      `json:"user_status"`
}
