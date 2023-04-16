package jwt_auth

type UserUsecase interface {
	FindByUsername(username string) (string, string)
}

type userUsecase struct {
	userRepo UserRepo
}

func (u userUsecase) FindByUsername(username string) (string, string) {

	return u.userRepo.GetByUsername(username)
}

func NewUserUsecase(userRepo UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
