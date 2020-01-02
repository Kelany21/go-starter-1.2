package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/config"
	"time"
)

type HashTag struct {
	gorm.Model
	HashTag          string    `gorm:"index:hash_tag" json:"hash_tag" binding:"required"`
	WordsCount 		 int       `gorm:"index:words_count" json:"words_count;unsigned"`
	GoodCount        int       `gorm:"index:good_count;unsigned" json:"good_count"`
	BadCount         int       `gorm:"index:bad_count;unsigned" json:"bad_count"`
	TweetCount       int       `gorm:"index:tweet_count;unsigned"  json:"tweet_count"`
	NeutralCount     int       `gorm:"index:neutral_count;unsigned" json:"neutral_count"`
	SetsCount        int       `gorm:"index:sets_count;unsigned" json:"sets_count"`
	BadWordCount     int       `gorm:"index:bad_word_count;unsigned" json:"bad_word_count"`
	GoodWordCount    int       `gorm:"index:good_word_count;unsigned" json:"good_word_count"`
	NeutralWordCount int       `gorm:"index:neutral_word_count;unsigned" json:"neutral_word_count"`
	UserId           int       `gorm:"index:user_id" json:"user_id"`
	WhoAdd           string    `gorm:"index:who_add" json:"who_add"`
	LatestSync       time.Time `gorm:"index:latest_sync" json:"latest_sync"`
	CountMoveGood    int       `gorm:"index:count_move_good;unsigned" json:"count_move_good"`
	CountMoveBad     int       `gorm:"index:count_move_bad;unsigned" json:"count_move_bad"`
	CountMoveNeutral int       `gorm:"index:count_move_neutral;unsigned" json:"count_move_neutral"`
	Stream           int       `gorm:"index:stream" json:"stream" binding:"required"`
	ActionId 		 int 	   `gorm:"index:action_id;default:2" json:"action_id" binding:"required"`
}

func (s *MigrationTables) HashTagMigrate() {
	config.DB.AutoMigrate(&HashTag{})
}


func HashTagFillAbleColumn()[]string {
	return []string{"hash_tag"}
}