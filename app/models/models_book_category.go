package models

type BookCategory struct {
	CustomGormModel
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type: text"`
	Status      bool   `json:"status" gorm:"type: bool"`
}
