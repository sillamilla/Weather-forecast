package user

type Service interface {
}

type userService struct {
}

func NewUserService() Service {
	return userService{}
}
