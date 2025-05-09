package entity

import "time"

type Customer struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	Gender    string     `json:"gender"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy *string    `json:"updated_by" db:"updated_by"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
	DeletedBy *string    `json:"-" db:"deleted_by"`

	Total int `json:"-" db:"total"`
}
