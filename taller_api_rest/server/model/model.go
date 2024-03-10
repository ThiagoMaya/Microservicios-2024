package model

import "gorm.io/gorm"

// User representa un usuario en la base de datos
type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique"`
	Password string
}
