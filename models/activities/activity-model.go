package models

type Activity struct {
	ActivityID  int    `gorm:"primary_key"`
	UserID      int    `gorm:"type:int;not null"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	IsDone      bool   `gorm:"type:boolean;not null"`
}

type Activities []Activity
