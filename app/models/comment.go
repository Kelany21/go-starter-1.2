package models

import (
	"golang-starter/config"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserId  uint   `gorm:"type:int" json:"user_id"`
	PostId  uint   `gorm:"type:int" json:"post_id"`
	Comment string `gorm:"type:text" json:"comment"`
	Status  int8   `gorm:"type:tinyint" json:"status"`
}

func (s *MigrationTables) CommentMigrate() {
	config.DB.AutoMigrate(&Comment{})
}

func CommentFillAbleColumn() []string {
	return []string{"comment", "status"}
}

func ActiveComment(db *gorm.DB) *gorm.DB {
	return db.Where("status = 2")
}
