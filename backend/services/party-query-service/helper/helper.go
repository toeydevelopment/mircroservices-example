package helper

import (
	"encoding/base64"
	"errors"
)

type IHelper interface {
	DecodeCursor(cursor string) (string, error)
	EncodeCursor(lastID string) string
}

type Helper struct{}

func New() IHelper {
	return Helper{}
}

func (Helper) EncodeCursor(lastID string) string {

	reverseLastID := reverse(lastID)

	return base64.RawStdEncoding.EncodeToString([]byte(reverseLastID))
}

func (Helper) DecodeCursor(cursor string) (string, error) {

	decoded, err := base64.RawStdEncoding.DecodeString(cursor)

	if err != nil {
		return "", errors.New("incorrect cursor")
	}

	lastID := reverse(string(decoded))

	return lastID, nil
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
