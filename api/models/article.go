package models

import "time"

type Article struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Author      string    `json:"author"`
    PublishDate time.Time `json:"publish_date"`
    Title       string    `json:"title"`
    Subtitle    string    `json:"subtitle"`
    MainImage   int       `json:"main_image" gorm:"foreignKey:ID;references:media_id"`
    Content     string    `json:"content"`
}
