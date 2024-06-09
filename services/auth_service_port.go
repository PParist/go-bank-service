package service

type AuthService interface {
	UserLogin(string, string) (string, error)
}
