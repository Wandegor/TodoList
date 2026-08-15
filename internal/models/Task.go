package models

type Task struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Text      string `gorm:"size:1000;not null" json:"text"`
	Completed bool   `gorm:"not null;default:false" json:"completed"`
}
