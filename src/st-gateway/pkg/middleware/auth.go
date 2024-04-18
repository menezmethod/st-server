package middleware

import (
	"net/http"
	"strings"

	pbAuth "github.com/menezmethod/st-server/src/st-gateway/pkg/pb/auth"
)

// AuthMiddleware creates a middleware to handle authentication using pb.AuthServiceClient.
func AuthMiddleware(authClient pbAuth.AuthServiceClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			skipAuthPaths := map[string]bool{
				"/v1/auth/login":    true,
				"/v1/auth/register": true,
			}

			if _, ok := skipAuthPaths[r.URL.Path]; ok {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing authorization header", http.StatusUnauthorized)
				return
			}

			authHeaderParts := strings.Split(authHeader, "Bearer ")
			if len(authHeaderParts) != 2 {
				http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
				return
			}

			token := authHeaderParts[1]
			_, err := authClient.Validate(r.Context(), &pbAuth.ValidateRequest{Token: token})
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
