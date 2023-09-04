package models

type TextContent struct {
    ContentID   uint   `json:"content_id" gorm:"primaryKey"`
    ContentType string `json:"content_type"`
    ContentText string `json:"content_text"`
}