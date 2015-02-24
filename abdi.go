package abdi

import (
	"crypto/hmac"
	"crypto/sha256"
	"unicode/utf8"

	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordTooShort    = errors.New("passwords must be 8 characters or greater")
	ErrBlacklistedPassword = errors.New("password is on the blacklist")
)

// Key is the nonce used for HMAC
// From Mozilla - " The nonce for the hmac value is designed to be stored on the
// file system and not in the databases storing the password hashes. In the
// event of a compromise of hash values due to SQL injection, the nonce will
// still be an unknown value since it would not be compromised from the file
// system. This significantly increases the complexity of brute forcing the
// compromised hashes considering both bcrypt and a large unknown nonce value"
var Key []byte

// Cost used for the bcrypt hash
// minimum is 4, maximum is 31
var Cost = 10

// MinPasswordLength allowed by abdi
var MinPasswordLength = 8

// Hash the password first using salted bcrypt, then sign the hash using HMAC
func Hash(password string) (string, error) {
	if utf8.RuneCountInString(password) < MinPasswordLength {
		return "", ErrPasswordTooShort
	}

	if err := checkBlacklist(password); err != nil {
		return "", err
	}

	passbyte := []byte(password)
	// HMAC the password
	signedPass := sign(passbyte)
	// salt + bcrypt the signed password
	hashedPass, err := hash(signedPass)
	if err != nil {
		return "", err
	}

	// encode as base64 to return
	encoded := base64.StdEncoding.EncodeToString(hashedPass)
	return encoded, nil
}

// Check the password against a hash
func Check(password string, oldHash string) error {
	hashed, err := base64.StdEncoding.DecodeString(oldHash)
	if err != nil {
		return err
	}

	signedPass := sign([]byte(password))

	err = bcrypt.CompareHashAndPassword(hashed, signedPass)
	if err != nil {
		return err
	}

	return nil
}

func checkBlacklist(pass string) error {
	for _, str := range Blacklist {
		if pass == str {
			return ErrBlacklistedPassword
		}
	}
	return nil
}

func hash(pass []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(pass, Cost)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func sign(pass []byte) []byte {
	mac := hmac.New(sha256.New, Key)
	mac.Write(pass)
	expectedMAC := mac.Sum(nil)
	return expectedMAC
}
