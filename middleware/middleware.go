package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/controller"
	"github.com/Mau005/MyExpenses/models"
)

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, thorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := configuration.Store.Get(r, configuration.NAME_SESSION)
		if err != nil {
			http.Error(w, "Error al obtener la sesión", http.StatusInternalServerError)
			return
		}
		var api controller.ApiController
		if tokenStr, ok := session.Values["token"].(string); !ok {
			w.WriteHeader(http.StatusNetworkAuthenticationRequired)
			json.NewEncoder(w).Encode(models.Exception{
				Error:         configuration.ERROR_SERVICE_USER,
				Status:        http.StatusNetworkAuthenticationRequired,
				Message:       configuration.ERROR_PRIVILEGES_GEN,
				TimeStamp:     time.Now(),
				TransactionId: "1",
				CorrelationId: "1",
			})
			return
		} else {
			err = api.AuthenticateJWT(tokenStr)
			if err != nil {
				w.WriteHeader(http.StatusNetworkAuthenticationRequired)
				json.NewEncoder(w).Encode(models.Exception{
					Error:         configuration.ERROR_SERVICE_USER,
					Status:        http.StatusNetworkAuthenticationRequired,
					Message:       configuration.ERROR_PRIVILEGES_GEN,
					TimeStamp:     time.Now(),
					TransactionId: "1",
					CorrelationId: "1",
				})
				return
			}

		}

		next.ServeHTTP(w, r)

	})
}
