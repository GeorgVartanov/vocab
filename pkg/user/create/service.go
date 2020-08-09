package create

//ServiceCreater
type ServiceCreater interface {
	Create ( User) (error)
}

type service struct {
	RepositoryCreater
}

func NewService(r RepositoryCreater) ServiceCreater {
	return &service{r}
}

func (s *service) Create(user User) ( error){
	if err := s.RepositoryCreater.Create(user); err!=nil{
		return  err
	}
	return nil
}
