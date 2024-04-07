package user

type ServiceInterface interface {
	Create(user *CreateUserRequest) error
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(user *CreateUserRequest) error {
	model := &Model{
		Age:     user.Age,
		Name:    user.Name,
		PetName: user.PetName,
	}
	return s.repo.Create(model)
}
