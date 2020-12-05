package model

import (
	"time"
)

//User struct
type User struct {
	ID        int64     `json:"id,string,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
