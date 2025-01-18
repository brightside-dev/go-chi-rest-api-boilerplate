package middlewares

import (
	"net/http"
	"time"

	customErr "github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/errors"
	"github.com/brightside-dev/go-chi-rest-api-boilerplate/internal/utils"

	"github.com/go-chi/jwtauth/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if token exists
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil || claims == nil {
			utils.WriteAPIErrorResponse(w, r, customErr.ErrInvalidJWTToken)
			return
		}

		// Verify the token's "exp" exists and is a string
		expClaim, ok := claims["exp"].(time.Time)
		if !ok {
			utils.WriteAPIErrorResponse(w, r, customErr.ErrInvalidJWTToken)
			return
		}

		// Check if the token has expired
		if time.Now().After(expClaim) {
			utils.WriteAPIErrorResponse(w, r, customErr.ErrJWTTokenExpired)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
