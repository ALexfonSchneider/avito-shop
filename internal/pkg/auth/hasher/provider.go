package hasher

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct{}

func (h BcryptHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h BcryptHasher) Compare(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func New() *BcryptHasher {
	return &BcryptHasher{}
}
