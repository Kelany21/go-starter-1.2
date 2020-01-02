package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/config"
)

type Set struct {
	gorm.Model
	Name       string `json:"name" binding:"required"`
	UseCount   int    `json:"use_count;unsigned"`
	WordsCount int    `json:"words_count;unsigned"`
	UserId     int    `json:"user_id"`
	WhoAdd     string `json:"who_add"`
	UsedNumber int    `json:"used_number"`
	TrashedBy  string `json:"trashed_by"`
	ActionId   int 	  `gorm:"default:12" json:"action_id" binding:"required"`
}

func (s *MigrationTables) SetMigrate() {
	config.DB.AutoMigrate(&Set{})
}

/**
* you can update these column only
 */
func SetFillAbleColumn() []string {
	return []string{"name"}
}
