package controller

import (
	"errors"
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
	tokenString, err = token.SignedString([]byte(configuration.Security))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func (ac *ApiController) SaveSession(tokenString *string, w http.ResponseWriter, r *http.Request) {
	session, _ := configuration.Store.Get(r, configuration.NAME_SESSION)
	if tokenString == nil {
		session.Values["token"] = nil
	} else {
		session.Values["token"] = *tokenString
	}

	session.Save(r, w)
}

func (ac *ApiController) AuthenticateJWT(tokenSession string) error {

	token, err := jwt.Parse(tokenSession, func(token *jwt.Token) (interface{}, error) {
		return []byte(configuration.Security), nil
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

func (ac *ApiController) GetSessionClaims(r *http.Request) (*models.Claims, error) {
	claims := &models.Claims{}
	session, err := configuration.Store.Get(r, configuration.NAME_SESSION)
	if err != nil {
		return claims, err
	}

	token, ok := session.Values["token"].(string)
	if !ok {
		return claims, errors.New("Token de session invalido")
	}
	tokenKey := []byte(configuration.Security)
	tokenParser := jwt.Parser{}

	_, err = tokenParser.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil {
		return claims, err
	}
	return claims, nil
}

func (ac *ApiController) GetSessionUser(r *http.Request) (models.User, error) {
	claims, err := ac.GetSessionClaims(r)
	if err != nil {
		return models.User{}, err
	}

	var uc UserController
	user, err := uc.GetUser(claims.UserName)
	if err != nil {
		return user, err
	}

	return user, nil
}
