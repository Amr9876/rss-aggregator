package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no auth info was found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("invalid auth info")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid first part of auth info")
	}

	return vals[1], nil

}
