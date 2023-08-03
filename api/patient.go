package api

type RegisterRequest struct {
	Username  string `json:"username" db:"username" binding:"required"`
	FirstName string `json:"firstname" db:"firstname" binding:"required"`
	LastName  string `json:"lastname" db:"lastname" binding:"required"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"password" db:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
