package usercase

type UserRepo interface {
	ChangeUserSegment(int, []string, []string) error
	GetUserSegment(int) ([]string, error)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ChangeUserSegment(id int, add []string, delete []string) error {
	return s.repo.ChangeUserSegment(id, add, delete)
}

func (s *UserService) GetUserSegment(id int) ([]string, error) {
	return s.repo.GetUserSegment(id)
}
