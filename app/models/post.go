package models

import (
	"golang-starter/config"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title      string `gorm:"type:varchar(50);" json:"title"`
	Body       string `gorm:"type:text;" json:"body"`
	Image      string `gorm:"type:varchar(255);" json:"image"`
	UserId     uint   `gorm:"type:int" json:"user_id"`
	CategoryId uint   `gorm:"type:int" json:"category_id"`
	Status     int8   `gorm:"type:tinyint" json:"status"`
}

func (s *MigrationTables) PostMigrate() {
	config.DB.AutoMigrate(&Post{})
}

func PostFillAbleColumn() []string {
	return []string{"title", "body", "status", "image"}
}

func ActivePost(db *gorm.DB) *gorm.DB {
	return db.Where("status = 2")
}
