package models

type Book struct {
	CustomGormModel
	Title           string `json:"title" gorm:"type: varchar(255)"`
	Author          string `json:"author" gorm:"type: varchar(255)"`
	BookCode        string `json:"book_code" gorm:"type: varchar(20)"`
	PublicationYear int    `json:"publication_year" gorm:"type:int"`
	Stock           int    `json:"stock" gorm:"type:int"`
	Language        string `json:"language" gorm:"type: varchar(50)"`
	Publisher       string `json:"publisher" gorm:"type: varchar(50)"`
	Description     string `json:"description" gorm:"type: text"`
	NumberOfPages   int    `json:"number_of_pages" gorm:"type: int"`
	CategoryID      int    `json:"category" gorm:"type: int"`
	Image           string `json:"image" gorm:"type: text"`
	Status          bool   `json:"status" gorm:"type: bool"`
}
