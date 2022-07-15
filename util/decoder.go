package util

import (
	"encoding/base64"
	"errors"
)

func DecodeBase64String(secret string) (string, error) {
	bytes, err := decodeSecretKey(secret)
	if err != nil {
		return "", err
	}

	bytes, err = ReverseBytes(bytes)
	if err != nil {
		return "", err
	}

	bytes, err = ReplaceBytes(bytes, ':', ' ')
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ReplaceBytes(bytes []byte, target int32, char int32) ([]byte, error) {
	if bytes == nil {
		return nil, errors.New("bytes is nil")
	}

	result := make([]byte, len(bytes))
	for index, value := range bytes {
		b := value
		if value == byte(target) {
			b = byte(char)
		}
		result[index] = b
	}

	return result, nil
}

func decodeSecretKey(secret string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func ReverseBytes(bytes []byte) ([]byte, error) {
	if bytes == nil {
		return nil, errors.New("bytes is nil")
	}

	result := make([]byte, len(bytes))

	i := 0
	for j := len(bytes) - 1; j >= 0; j-- {
		result[i] = bytes[j]
		i++
	}

	return result, nil
}
