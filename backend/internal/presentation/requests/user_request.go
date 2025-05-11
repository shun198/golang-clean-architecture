package requests

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin general"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin general"`
}

type ListUsersQuery struct {
	Limit    int    `form:"limit" binding:"gte=0,max=100"`
	Offset   int    `form:"offset" binding:"gte=0"`
	Email    string `form:"email" binding:"omitempty"`
	Username string `form:"username" binding:"omitempty"`
	Role     string `form:"role" binding:"omitempty"`
}
