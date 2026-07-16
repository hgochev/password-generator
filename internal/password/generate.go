package password

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/hgochev/password-generator/internal/models"
)

const (
	lower   = "abcdefghijklmnopqrstuvwxyz"
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	special = "!@#$%^&*()_+><~{}[]"
)

func Generate(flags models.Options) (string, error) {
	if flags.Length <= 0 {
		return "", errors.New("password length must be greater than 0")
	}

	// Build the character pool and guarantee at least one char from each enabled set.
	// Note: NoNumbers/NoSpecialChars/NoMixedCase are false when the user passed -nn/-ns/-nm (i.e. disabled).
	pool := lower
	guaranteed := []byte{}

	ch, err := randomChar(lower)
	if err != nil {
		return "", err
	}
	guaranteed = append(guaranteed, ch)

	if !flags.NoMixedCase {
		pool += upper
		if ch, err = randomChar(upper); err != nil {
			return "", err
		}
		guaranteed = append(guaranteed, ch)
	}

	if !flags.NoNumbers {
		pool += digits
		if ch, err = randomChar(digits); err != nil {
			return "", err
		}
		guaranteed = append(guaranteed, ch)
	}

	if !flags.NoSpecialChars {
		pool += special
		if ch, err = randomChar(special); err != nil {
			return "", err
		}
		guaranteed = append(guaranteed, ch)
	}

	if flags.Length < len(guaranteed) {
		return "", fmt.Errorf("length %d is too short for the selected character sets (minimum %d)", flags.Length, len(guaranteed))
	}

	password := make([]byte, flags.Length)
	copy(password, guaranteed)

	for i := len(guaranteed); i < flags.Length; i++ {
		if password[i], err = randomChar(pool); err != nil {
			return "", err
		}
	}

	// Shuffle so guaranteed chars aren't always at the start.
	for i := len(password) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", err
		}
		password[i], password[j.Int64()] = password[j.Int64()], password[i]
	}

	return string(password), nil
}

func randomChar(charset string) (byte, error) {
	idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return 0, err
	}
	return charset[idx.Int64()], nil
}
