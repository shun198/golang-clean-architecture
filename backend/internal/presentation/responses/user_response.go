package responses

import "time"

type UsersResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedBy int       `json:"updated_by"`
}

type ListUsersResponse struct {
	Count  int64           `json:"count"`
	Length int             `json:"length"`
	Users  []UsersResponse `json:"users"`
}
