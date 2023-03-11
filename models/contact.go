package models

import "gorm.io/gorm"

// people and people relationSHIP
type Contact struct {
	gorm.Model
	// people and people contact
	OwnerId  uint
	TargetId uint
	// 0 , 1 ,2
	Type int
	Desc string
}

// define mysql table name
func (contact *Contact) TableName() string {
	return "contact"
}
