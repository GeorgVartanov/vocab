package read

//ServiceReader
type ServiceReader interface {
	Read(int) (User, error)
	ReadAll() ([]User, error)
}

type service struct {
	RepositoryReader
}

func NewService(r RepositoryReader) ServiceReader {
	return &service{r}
}

func (s *service) Read(id int) (User, error) {
	user, err := s.RepositoryReader.Read(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) ReadAll() ([]User, error) {
	users, err := s.RepositoryReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
