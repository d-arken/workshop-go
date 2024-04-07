package user

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	PetName string `json:"pet_name,omitempty"`
}

func (m Model) TableName() string {
	return "users"
}
