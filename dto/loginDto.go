package dto

type LoginDTO struct {
	Email    string `gorm:"email" form:"email" binding:"required" validate:"email"`
	Password string `gorm:"password" form:"password" binding:"required" validate:"min:6"`
}
