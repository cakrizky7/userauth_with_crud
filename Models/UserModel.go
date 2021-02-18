package Models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Level    string    `json:"level"`
	Company  string    `json:"company"`
}

func (b *User) TableName() string {
	return "user"
}
