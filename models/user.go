package models

import (
	"time"
)

type User struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name"`
    Email     string    `json:"email" gorm:"unique;not null"`           
    Password  string    `json:"-" gorm:"not null"`                     
    CreatedAt time.Time `json:"created_at"`  
}