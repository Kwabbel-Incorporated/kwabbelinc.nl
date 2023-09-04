package models

type Media struct {
    MediaID   int    `json:"media_id" gorm:"primaryKey"`
    Filename  string `json:"filename"`
    FileType  string `json:"file_type"`
    FilePath  string `json:"file_path"`
}