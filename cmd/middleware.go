package main

import (
	"errors"
	"net/http"

	"github.com/brightside-dev/go-chi-rest-api-boilerplate/cmd/utils"

	"github.com/go-chi/jwtauth/v5"
)

// CustomAuthenticator is a middleware that customizes the error response for JWT authentication
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if token exists and is valid
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil || claims == nil {
			// Customize the error response
			customErr := errors.New("invalid token")
			utils.WriteAPIErrorResponse(w, r, customErr)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
