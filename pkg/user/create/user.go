package create

type User struct {
	FirstName string `json:"firstName" db:"firstName"`
	LastName  string `json:"lastName" db:"lastName"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

// RepositoryCreater ...
type RepositoryCreater interface {
	Create(User) error
}
