package login

import "github.com/georgvartanov/vocabProject/pkg/user/storage"

//ServiceReader
type ServiceLoginer interface {
	GetUserByEmail(email string) (storage.User, error)
}

type service struct {
	RepositoryLoging
}

func NewService(r RepositoryLoging) ServiceLoginer {
	return &service{r}
}

func (s *service) GetUserByEmail(email string) (storage.User, error) {
	user, err := s.RepositoryLoging.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}


