package models

import (
    "time"
)

type User struct {
    Id              uint            `json:"id" gorm:"primaryKey"`
    Name            string          `json:"name,omitempty" validate:"required"`
    Email           string          `json:"email,omitempty" validate:"required"`
    UserRoleId      uint            `json:"user_role_id,omitempty" validate:"required"`
    CreatedAt       time.Time
    UpdatedAt       time.Time
}