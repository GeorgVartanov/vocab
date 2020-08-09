package login

import (
	"github.com/georgvartanov/vocabProject/pkg/user/storage"
	"time"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"firstName"`
	LastName  string `json:"lastName" db:"lastName"`
	Email     string `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Created time.Time `json:"created" db:"created"`
	Changed time.Time `json:"changed" db:"changed"`
}

type RepositoryLoging interface {
	GetUserByEmail(email string) (storage.User, error)

}