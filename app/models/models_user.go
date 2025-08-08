package models

import (
	"time"

	"github.com/guregu/null"
	"gorm.io/gorm"
)

type GlobalUser struct {
	CustomGormModel
	Avatar          string    `json:"avatar" gorm:"type: text"`
	Fullname        string    `json:"fullname" gorm:"type: varchar(255)"`
	Email           string    `json:"email" gorm:"type: varchar(255)"`
	IdAnggota       string    `json:"id_anggota" gorm:"type: varchar(255)"`
	Password        string    `json:"-" gorm:"type: varchar(255)"`
	Phone           string    `json:"phone" gorm:"type: varchar(50)"`
	Address         string    `json:"address" gorm:"type: varchar(255)"`
	Village         string    `json:"village" gorm:"type: varchar(255)"`
	District        string    `json:"district" gorm:"type: varchar(255)"`
	City            string    `json:"city" gorm:"type: varchar(255)"`
	Province        string    `json:"province" gorm:"type: varchar(255)"`
	Country         string    `json:"country"`
	ZipCode         string    `json:"zip_code" gorm:"type: varchar(50)"`
	Status          int       `json:"status" gorm:"type: int4"`
	RoleID          int       `json:"role_id" gorm:"type: int8"`
	EmailVerifiedAt null.Time `json:"email_verified_at" gorm:"type: timestamptz"`
	Gender          string    `json:"gender" gorm:"column:gender"`
}

func (GlobalUser) TableName() string {
	return "global_users"
}

type CustomGormModel struct {
	ID        uint            `gorm:"primary_key" json:"id"`
	EncodedID string          `json:"encoded_id" gorm:"-"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
