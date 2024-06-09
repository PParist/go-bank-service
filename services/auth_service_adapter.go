package service

import (
	"os"
	"time"

	"github.com/PParist/go-bank-service/errorhandler"
	"github.com/PParist/go-bank-service/logs"
	"github.com/PParist/go-bank-service/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) UserLogin(username string, password string) (string, error) {

	user, err := s.repo.Login(username)
	if err != nil {
		logs.Error(err)
		if err == gorm.ErrRecordNotFound {
			return "", errorhandler.NewErrorNotFound("user not found")
		}
		return "", errorhandler.NewErrorInternalServerError(err.Error())
	}

	//TODO: compare password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logs.Error(err)
		return "", errorhandler.NewErrorUnauthorized("Unauthorized")
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"iss":      "myapp",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
		"iat":      time.Now().Unix(),
		"user_uid": user.User_uid,
		"role":     user.User_Role,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRECT")))
	if err != nil {
		logs.Error(err)
		return "", errorhandler.NewErrorInternalServerError("can't signed string jwt")
	}
	return t, nil
}
