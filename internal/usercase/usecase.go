package usercase

type Repo interface {
	UserRepo
	SegmentRepo
}

type UseCase struct {
	*UserService
	*SegmentService
}

func NewUseCase(repo Repo) *UseCase {
	return &UseCase{UserService: NewUserService(repo), SegmentService: NewSegmentService(repo)}
}
