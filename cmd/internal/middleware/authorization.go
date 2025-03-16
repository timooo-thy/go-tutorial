package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/timooo-thy/go_tutorial/api"
	"github.com/timooo-thy/go_tutorial/internal/tools"
)

var UnauthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.URL.Query().Get("Authorization")
		var err error
		
		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.loginDetails
		loginDetails, err = database.GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).Token) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}