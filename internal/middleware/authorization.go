package middleware

import (
	"errors"
	"net/http"

	"github.com/Yeet94/Go-try/api"
	"github.com/Yeet94/Go-try/internal/tools"
	log "github.com/sirupsen/logrus"
)

var ErrUnauthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		log.Infof("Validating username: %s, token: %s", username, token) // Log username and token for debugging
		if username == "" || token == "" {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			log.Errorf("Failed to initialize database: %v", err)
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := (*database).GetUserLoginDetails(username)

		if loginDetails == nil {
			log.Errorf("No login details found for username: %s", username)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		log.Infof("Retrieved login details: %+v", loginDetails) // Log retrieved login details for debugging

		if token != loginDetails.AuthToken {
			log.Errorf("Token mismatch for username: %s. Provided: %s, Expected: %s", username, token, loginDetails.AuthToken)
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
