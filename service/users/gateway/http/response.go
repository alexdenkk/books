package http

// AuthResponse - response struct for JWT token
type AuthResponse struct {
	Token string `json:"token"`
}
