package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/utils"

	"github.com/go-chi/jwtauth/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if token exists
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil || claims == nil {
			// Customize the error response
			customErr := errors.New("invalid token")
			utils.WriteAPIErrorResponse(w, r, customErr)
			return
		}

		// Verify the token's "exp" exists and is a float64
		expClaim, ok := claims["exp"].(float64) // The "exp" is usually a float64
		if !ok {
			customErr := errors.New("invalid token: missing or malformed 'exp' claim")
			utils.WriteAPIErrorResponse(w, r, customErr)
			return
		}

		// Check if the token has expired
		if time.Now().Unix() > int64(expClaim) {
			customErr := errors.New("token has expired")
			utils.WriteAPIErrorResponse(w, r, customErr)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
