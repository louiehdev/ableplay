package auth

// Look at https://docs.sqlc.dev/en/latest/howto/transactions.html for how to do sql transactions for submitting changes and updating the change status at the same time

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/alexedwards/argon2id"
)

const (
	RoleUser      = "user"
	RoleModerator = "moderator"
	RoleAdmin     = "admin"
)

var RoleLevel = map[string]int{
	RoleUser:      1,
	RoleModerator: 2,
	RoleAdmin:     3,
}

func CreateHash(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func CheckHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return match, nil
}

func GetAPIKey(r *http.Request) (rawKey string, prefixKey string, err error) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		apiKey = r.URL.Query().Get("api_key")
	}

	if apiKey == "" {
		return "", "", fmt.Errorf("missing API key")
	}

	return apiKey, strings.TrimPrefix(apiKey, "ablply_")[:8], nil
}

func CreateAPIKey() (rawKey string, prefixKey string) {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	key := base64.RawURLEncoding.EncodeToString(bytes)
	return "ablply_" + key, key[:8]
}
