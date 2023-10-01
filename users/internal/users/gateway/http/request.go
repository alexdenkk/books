package http

// AuthRequest - request struct for authorization
type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// ChangePasswordRequest - request struct for changing password
type ChangePasswordRequest struct {
	Old string `json:"old"`
	New string `json:"new"`
}
