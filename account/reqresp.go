package account

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}
	GetUserRequest struct {
		Id string `json:"id"`
	}
)
