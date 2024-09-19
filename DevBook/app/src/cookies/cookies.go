package cookies

import (
	"devbook_app/src/config"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

const CookieName = "session"

var secureCookie *securecookie.SecureCookie

func Configure() {
	secureCookie = securecookie.New(
		config.HashKey,
		config.BlockKey,
	)
}

func Save(w http.ResponseWriter, ID, token string, username string) error {
	data := map[string]string{
		"ID":       ID,
		"token":    token,
		"username": username,
	}

	encoded, err := secureCookie.Encode(CookieName, data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    encoded,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(6 * time.Hour),
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return nil, err
	}

	cookie.Valid()

	data := make(map[string]string)
	if err = secureCookie.Decode(CookieName, cookie.Value, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func IsValid(r *http.Request) bool {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return false
	}

	if err = cookie.Valid(); err != nil {
		return false
	}

	if cookie.Expires.Before(time.Now()) {
		return false
	}

	return true
}

func Clear(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   0,
		Expires:  time.Unix(0, 0),
	})
}

func GenerateKeyPair() ([]byte, []byte) {
	hashKey := securecookie.GenerateRandomKey(32)
	blockKey := securecookie.GenerateRandomKey(16)

	return hashKey, blockKey
}
