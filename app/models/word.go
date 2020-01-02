package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/config"
)

type Word struct {
	gorm.Model
	Type       string `gorm:"index:type" json:"type"`
	Word       string `gorm:"index:word" json:"word" binding:"required"`
	SourceId   int    `gorm:"index:source_id" json:"source_id" binding:"required"`
	SourceType string `gorm:"index:source_type" json:"source_type" binding:"required"`
}

func (s *MigrationTables) WordMigrate() {
	config.DB.AutoMigrate(&Word{})
}

/**
* you can update these column only
 */
func WordFillAbleColumn() []string {
	return []string{"type", "word"}
}
