package service

import (
	"github.com/mashingan/smapping"
	"github.com/viky1sr/latihan_go-lang.git/dto"
	"github.com/viky1sr/latihan_go-lang.git/entity"
	"github.com/viky1sr/latihan_go-lang.git/repository"
	"log"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	FindByProfile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (u userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	return u.userRepository.UpdateUser(userToUpdate)
}

func (u userService) FindByProfile(userID string) entity.User {
	return u.userRepository.ProfileUser(userID)
}
