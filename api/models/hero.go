package models

type Hero struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Title         string `json:"title"`
	Image         uint   `json:"image" gorm:"foreignKey:ID;references:media_id"`
	Text          string `json:"text"`
	ButtonText    string `json:"button_text"`
	ButtonURL     string `json:"button_url"`
	IsCurrentHero uint   `json:"is_current_hero"`
}