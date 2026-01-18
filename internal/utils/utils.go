package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
)

func ValidatePhoneNumber(phone string) bool {

	if len(phone) != 11 {
		return false
	}

	re := regexp.MustCompile(`^1[3-9]\d{9}$`)

	if re.MatchString(phone) {
		return true
	}
	return false
}

func GenerateCode() (string, error) {

	max := big.NewInt(1000000) // 0 ~ 999999

	n, err := rand.Int(rand.Reader, max)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%06d", n.Int64()), nil
}
