package models

import (
	"github.com/guregu/null"
)

type BookLending struct {
	CustomGormModel
	BookID     int       `json:"book_id" gorm:"type: int;not null"`
	UserID     int       `json:"user_id" gorm:"type: int;not null"`
	BorrowDate null.Time `json:"borrow_date"  gorm:"type:timestamptz"`
	DueDate    null.Time `json:"due_date"  gorm:"type:timestamptz"`
	ReturnDate null.Time `json:"return_date"  gorm:"type:timestamptz"`
	Status     string    `json:"status" gorm:"type: varchar(50)"` // misalnya: "requested","approved", "borrowed","rejected" "returned", "overdue"
	Notes      string    `json:"notes" gorm:"type: text"`
}
