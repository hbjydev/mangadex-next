package models

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hbjydev/mangadex-next/database"
	"golang.org/x/crypto/bcrypt"
)

type password string
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  password  `json:"password"`
	Email     string    `json:"email"`
	LevelID   string    `json:"level_id"`
	JoinedAt  time.Time `json:"joined_at"`
	UpdateAt  time.Time `json:"update_at"`
	LastSeen  time.Time `json:"last_seen"`
	Website   string    `json:"website"`
	Biography string    `json:"biography"`
	Views     int64     `json:"views"`
	Uploads   int64     `json:"uploads"`
	Premium   bool      `json:"premium"`
	MDAtHome  bool      `json:"md_at_home"`
	AvatarURL string    `json:"avatar"`
}

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}

func UserByUsername(username string) (*User, error) {
	row := database.DB.QueryRow(`
        SELECT
            hex(id), username, email, password, level_id, last_seen, website,
            biography, views, uploads, premium, md_at_home, avatar_url,
            joined_at, update_at FROM users
        WHERE username = ?
    `, username)

	var u User
	var ls mysql.NullTime
	var ja mysql.NullTime
	var ua mysql.NullTime

	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.LevelID,
		&ls, &u.Website, &u.Biography, &u.Views, &u.Uploads, &u.Premium,
		&u.MDAtHome, &u.AvatarURL, &ja, &ua); err != nil {
		log.Printf("UserByUsername: error at SQL query: %v\n", err)
		return nil, err
	}

	if ls.Valid {
		u.LastSeen = ls.Time
	} else {
		log.Printf("UserByUsername: invalid SQL datetime value (id %v)\n", u.ID)
		return nil, errors.New("invalid sql datetime value")
	}

	if ja.Valid {
		u.JoinedAt = ja.Time
	} else {
		log.Printf("UserByUsername: invalid SQL datetime value (id %v)\n", u.ID)
		return nil, errors.New("invalid sql datetime value")
	}

	if ua.Valid {
		u.UpdateAt = ua.Time
	} else {
		log.Printf("UserByUsername: invalid SQL datetime value (id %v)\n", u.ID)
		return nil, errors.New("invalid sql datetime value")
	}

	return &u, nil
}

func UserFromJSON(jsonString string) (*User, error) {
	var user User

	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) CheckPassword(candidate string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(candidate))
	if err != nil {
		return false
	}
	return true
}

func (u *User) Normalize() (*string, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	stringVal := string(bytes)

	return &stringVal, nil
}
