package jwt

type JwtPayloadModel struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      uint   `json:"age"`
	IsActive bool   `json:"is_active"`
}
