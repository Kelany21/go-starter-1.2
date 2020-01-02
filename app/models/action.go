package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/config"
)

type Action struct {
	gorm.Model
	Title  string `gorm:"type:varchar(50)" json:"title"`
	Count	int `json:"count"`
	ModuleName string `json:"module_name"`

}

func (s *MigrationTables) ActionMigrate() {
	config.DB.AutoMigrate(&Action{})
}