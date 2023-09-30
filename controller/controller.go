package controller

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type ApiController struct{}

func (ac *ApiController) GenerateCryptPassword(password string) string {
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hasedPassword)
}

func (ac *ApiController) CompareCryptPassword(password, passwordTwo string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordTwo))
}

func (ac *ApiController) generateClaims(user models.User) *models.Claims {
	expirationTime := time.Now().Add(configuration.EXPIRATION_TOKEN * time.Hour)
	return &models.Claims{
		UserName: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
}

func (ac *ApiController) GenerateToken(user models.User) (tokenString string, err error) {

	claims := ac.generateClaims(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte("secret-key"))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func (ac *ApiController) SaveSession(tokenString *string, w http.ResponseWriter, r *http.Request) {
	session, _ := configuration.STORE.Get(r, configuration.NAME_SESSION)
	if tokenString == nil {
		session.Values["token"] = nil
	} else {
		session.Values["token"] = *tokenString
	}

	session.Save(r, w)
}

func (ac *ApiController) GenerateSecretKey(lenSecurity int) (string, error) {
	key := make([]byte, lenSecurity)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	encodedKey := base64.StdEncoding.EncodeToString(key)

	return encodedKey, nil
}

func (ac *ApiController) AuthenticateJWT(tokenSession string) error {

	token, err := jwt.Parse(tokenSession, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})

	if err != nil || !token.Valid {
		return err
	}

	return nil

}

// Method default ID for generate in api
func (ac *ApiController) ParseUintDefault(id string, defaultId uint) uint {
	idNew, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return defaultId
	}
	return uint(idNew)
}
