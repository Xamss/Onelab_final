package domain

type patient struct {
	ID         int64  `json:"id" db:"id"`
	Username   string `json:"username" db:"username" binding:"required"`
	FirstName  string `json:"first_name" db:"first_name" binding:"required"`
	LastName   string `json:"last_name" db:"last_name" binding:"required"`
	Speciality string `json:"speciality" db:"speciality" binding:"required"`
	Password   string `json:"password" db:"hashed_password" binding:"required"`
}
