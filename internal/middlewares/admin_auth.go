package middlewares

import (
	"errors"
	"net/http"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/utils"
)

func AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		customErr := errors.New("token has expired")
		utils.WriteAPIErrorResponse(w, r, customErr)
		return

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
