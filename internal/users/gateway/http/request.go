package http

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	Old string `json:"old"`
	New string `json:"new"`
}
