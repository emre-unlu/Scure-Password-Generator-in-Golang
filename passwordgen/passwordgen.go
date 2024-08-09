package passwordgen

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	specialChars = "!@#$%^&*()_+{}:\"<>?|[];',./`~"
	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      = "0123456789"
)

func GeneratePassword(length int) (string, error) {
	if length < 4 {
		return "", fmt.Errorf("password length must be at least 4")
	}

	allChars := []string{specialChars, lowercase, uppercase, numbers}
	password := make([]byte, length)

	// Step 1: Ensure one character from each category
	for i := 0; i < 4; i++ {
		char, err := pickCharFromCategory(allChars[i])
		if err != nil {
			return "", err
		}
		pos, err := randomInt(i, length-1)
		if err != nil {
			return "", err
		}
		password[i], password[pos] = password[pos], char
	}

	// Step 2: Fill the remaining characters
	for i := 4; i < length; i++ {
		categoryIndex, err := randomInt(0, 3)
		if err != nil {
			return "", err
		}
		char, err := pickCharFromCategory(allChars[categoryIndex])
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	return string(password), nil
}

func pickCharFromCategory(category string) (byte, error) {
	index, err := randomInt(0, len(category)-1)
	if err != nil {
		return 0, err
	}
	return category[index], nil
}

func randomInt(min, max int) (int, error) {
	if min > max || min < 0 {
		return 0, fmt.Errorf("invalid range (%d, %d)", min, max)
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + min, nil
}
