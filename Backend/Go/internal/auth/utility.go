package auth

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func sendEmail(to, subject, body string) error {
	fmt.Println("To:", to)
	fmt.Println("Subject:", subject)
	fmt.Println("Body:", body)
	return nil
}

func hashOTP(otp string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func generateAlphaNumericOTP(length int) (string, error) {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	otp := make([]byte, length)
	for i := 0; i < length; i++ {
		// get a secure random index
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		otp[i] = charset[num.Int64()]
	}

	return string(otp), nil
}
