package user

type Service interface {
}

type userService struct {
	us Service
}

func NewUserService(us Service) Service {
	return userService{us: us}
}
