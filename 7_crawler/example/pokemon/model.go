package pokemon

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name               string `json:"name"`
	Url                string `json:"url"`
	Weight             int    `json:"weight"`
	Height             int    `json:"height"`
	SpriteFrontDefault string `json:"sprite_front_default"`
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
}

type APIResponse struct {
	Name    string  `json:"name"`
	Url     string  `json:"url"`
	Weight  int     `json:"weight"`
	Height  int     `json:"height"`
	Sprites Sprites `json:"sprites"`
}

func (m *Model) TableName() string {
	return "pokemon"
}
