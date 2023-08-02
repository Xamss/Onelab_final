package domain

type User struct {
	ID        int64  `json:"id" db:"id"`
	Username  string `json:"username" db:"username" binding:"required"`
	FirstName string `json:"first_name" db:"firstname" binding:"required"`
	LastName  string `json:"last_name" db:"lastname" binding:"required"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"-" db:"hashed_password" binding:"required"`
}
