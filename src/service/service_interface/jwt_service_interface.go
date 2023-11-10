package service_interface

type JWTServiceInterface interface {
	GenerateToken(id uint) (string, error)
	GenerateHash(word string) (string, error)
	IsEqual(hash, word string) bool
}
