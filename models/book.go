package models

type Book struct {
	Id          int64  `grom:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(300);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Author      string `gorm:"type:varchar(100);not null" json:"author"`
	PublishDate string `gorm:"type:date;not null" json:"publish_date"`
}
