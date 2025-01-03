package service

type PasswordService interface {
	GenerateSalt(length int) (string, error)
	HashPassword(password, localSalt string) (string, error)
	VerifyPassword(plainPassword, hashedPassword, localSalt string) bool
}
