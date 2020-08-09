package read

import "time"

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"firstName"`
	LastName  string `json:"lastName" db:"lastName"`
	Email     string `json:"email" db:"email"`
	//Password  string    `json:"password" db:"password"`
	Created time.Time `json:"created" db:"created"`
	Changed time.Time `json:"changed" db:"changed"`
}

//RepositoryReader ...
type RepositoryReader interface {
	Read(int) (User, error)
	ReadAll() ([]User, error)
}

