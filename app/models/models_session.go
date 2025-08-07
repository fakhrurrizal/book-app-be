package models

type GlobalSignin struct {
	CustomGormModel
	UserID      int    `json:"user_id" gorm:"type: int8"`
	BearerToken string `json:"bearer_token" gorm:"type: text"`
	IPAddress   string `json:"ip_address" gorm:"type: varchar(255)"`
	UserAgent   string `json:"user_agent" gorm:"type: text"`
	HostName    string `json:"host_name" gorm:"type: varchar(255)"`
	AppID       int    `json:"app_id" gorm:"type: int8"`
	ISP         string `json:"isp" gomr:"type: varchar(255)"`
	City        string `json:"city" gomr:"type: varchar(255)"`
}

func (GlobalSignin) TableName() string {
	return "global_signins"
}
