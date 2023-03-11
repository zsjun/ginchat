package models

import "gorm.io/gorm"

// people and people relationSHIP
type GroupBacsic struct {
	gorm.Model
	// people and people contact
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string
}

// define mysql table name
func (group *GroupBacsic) TableName() string {
	return "group_basic"
}
