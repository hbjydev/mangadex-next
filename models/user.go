package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	LevelID   string    `json:"level_id"`
	JoinedAt  time.Time `json:"joined_at"`
	LastSeen  time.Time `json:"last_seen"`
	Website   string    `json:"website"`
	Biography string    `json:"biography"`
	Views     int64     `json:"views"`
	Uploads   int64     `json:"uploads"`
	Premium   bool      `json:"premium"`
	MDAtHome  bool      `json:"md_at_home"`
	AvatarURL string    `json:"avatar"`
}

func UserFromJSON(jsonString string) (*User, error) {
	var user User

	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Normalize() (*string, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	stringVal := string(bytes)

	return &stringVal, nil
}
