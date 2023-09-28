package configuration

import "github.com/gorilla/sessions"

var STORE *sessions.CookieStore

func GenerateSecretPassword() {
	STORE = sessions.NewCookieStore([]byte("secret-key"))
}
