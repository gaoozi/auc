package model

import "time"

type User struct {
	ID        int64 `gorm:"primary_key"`
	Nickname  string
	Username  string `gorm:"unique"`
	Password  string `gorm:"not null"`
	Avatar    string
	Gender    int8
	Phone     string
	Email     string    `gorm:"unique"`
	Status    int       `gorm:"default:0"`
	Role      int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
	IsDelete  int8      `gorm:"default:0"`
}

type RegisterRequest struct {
	Username      string
	Password      string
	CheckPassword string
}

