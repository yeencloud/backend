package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
}

func (User *User) Initialize() {
	User.ID = uuid.New()
}
