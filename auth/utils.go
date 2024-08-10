package auth

import (
	"encoding/base64"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

func NewRandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func StringToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64ToIDAndToken(s string) (uuid.UUID, string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return uuid.Nil, "", err
	}
	split := strings.Split(string(data), ":")
	id, err := uuid.Parse(split[0])
	if err != nil {
		return uuid.Nil, "", err
	}
	return id, split[1], nil
}

func Base64ToString(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
