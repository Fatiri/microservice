package models

//Account struct
type Account struct {
	AccountID uint64 `gorm:"primary_key; auto_increment; unique" json:"account_id"`
	Fullname  string `gorm:"not null; type:varchar(255);" json:"fullname,omitempty"`
	Place     string `gorm:"not null; type:varchar(255);" json:"place,omitempty"`
}
