package user

import (
	"api-go-hexa/business/user/model"
	"api-go-hexa/modules/user"
	"api-go-hexa/utils/jwt"
	"api-go-hexa/utils/password"
)

type Service interface {
	GetByID(id int) (*model.UserModel, error)
	UserRegister(um *model.UserModel) (*model.UserModel, error)
	UserLogin(um *model.UserLoginModel) (string, error)
	Update(id int, um *model.UserModel) error
	Delete(id int) error
}

type service struct {
	userRepository user.Repository
}

func NewUserService(u user.Repository) Service {
	return &service{
		userRepository: u,
	}
}

func (u *service) GetByID(id int) (*model.UserModel, error) {
	return u.userRepository.GetByID(id)
}

func (u *service) UserRegister(um *model.UserModel) (*model.UserModel, error) {
	um.Password, _ = password.HashPassword(um.Password)
	res, err := u.userRepository.UserRegister(um)
	if err != nil {
		return um, err
	}
	return res, nil
}

func (u *service) UserLogin(um *model.UserLoginModel) (string, error) {
	userObj, err := u.userRepository.UserLogin(um)
	if err != nil {
		return "", err
	}
	loggedIn, err := password.VerifyPassword(um.Password, userObj.Password)
	if err != nil || !loggedIn {
		return "", err
	}
	token, err := jwt.CreateJWTToken(userObj)
	return token, nil
}

func (u *service) Update(id int, um *model.UserModel) error {
	return u.userRepository.Update(id, um)
}

func (u *service) Delete(id int) error {
	return u.userRepository.Delete(id)
}
