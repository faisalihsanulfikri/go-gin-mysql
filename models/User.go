package models

import (
    "time"
)

type User struct {
    ID              uint           `gorm:"primaryKey"`
    Name            string
    Email           *string
    UserRoleId      uint
    CreatedAt       time.Time
    UpdatedAt       time.Time
}