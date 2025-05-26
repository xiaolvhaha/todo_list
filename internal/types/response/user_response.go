package response

type UserInfo struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	User  UserInfo `json:"user"`
	Token string   `json:"token"`
}
