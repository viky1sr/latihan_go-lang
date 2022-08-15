package service

import (
	"github.com/mashingan/smapping"
	"github.com/viky1sr/latihan_go-lang.git/dto"
	"github.com/viky1sr/latihan_go-lang.git/entity"
	"github.com/viky1sr/latihan_go-lang.git/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (a authService) VerifyCredential(email string, password string) interface{} {
	res := a.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (a authService) CreateUser(user dto.UserCreateDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := a.userRepository.InsertUser(userToCreate)
	return res
}

func (a authService) FindByEmail(email string) entity.User {
	return a.userRepository.FindByEmail(email)
}

func (a authService) IsDuplicateEmail(email string) bool {
	res := a.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
