package security

import "golang.org/x/crypto/bcrypt"

func HashString(plainString string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainString), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func CompareHashAndString(hashedString, plainString string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(plainString))
}
