package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// MD5 returns a md5 hash of src or a error.
func MD5(src string) (string, error) {

	file, err := os.Open(src)
	if err != nil {
		return "", err
	}

	defer file.Close()

	check := md5.New()

	if _, err := io.Copy(check, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(check.Sum(nil)), nil

}

// SHA256 returns a sha256 hash string from src or a error.
func SHA256(src string) (string, error) {

	file, err := os.Open(src)
	if err != nil {
		return "", err
	}

	defer file.Close()

	check := sha256.New()

	if _, err := io.Copy(check, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(check.Sum(nil)), nil

}
