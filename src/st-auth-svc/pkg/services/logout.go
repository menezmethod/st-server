package services

import (
	"errors"
	"net/http"
	"strings"
)

// LogoutHandler invalidates the user's token.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	token, err := extractTokenFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Example database call to blacklist the token
	err = Database.AddToTokenBlacklist(token)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	// Clear session cookie or perform other cleanup if necessary
	// For example, if you're using a session cookie:
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Database represents the database with a method to add tokens to the blacklist.
var Database = struct {
	AddToTokenBlacklist func(token string) error
}{
	AddToTokenBlacklist: func(token string) error {
		// Implement this function to add the token to the database blacklist
		return nil
	},
}

// extractTokenFromRequest retrieves the token from the Authorization header of the request.
func extractTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	// Assuming the Authorization header contains a bearer token.
	parts := strings.Split(authHeader, "Bearer ")
	if len(parts) != 2 {
		return "", errors.New("bearer token not provided")
	}

	return parts[1], nil
}
