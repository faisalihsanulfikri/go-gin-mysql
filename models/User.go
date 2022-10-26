package models

import (
    "time"
)

type User struct {
    Id              uint           `gorm:"primaryKey"`
    Name            string
    Email           *string
    UserRoleId      uint
    CreatedAt       time.Time
    UpdatedAt       time.Time
}