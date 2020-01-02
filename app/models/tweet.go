package models

import (
	"github.com/jinzhu/gorm"
	"golang-starter/config"
)

type Tweet struct {
	gorm.Model
	Text         string `gorm:"size:350" json:"text"`
	TweetID      int64  `json:"tweet_id"`
	Lang         string `json:"lang"`
	RetweetCount int    `json:"retweet_count"`
	ReplyCount   int    `json:"reply_count"`
	Place        string `json:"place"`
	CountryCode  string `json:"country_code"`
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	Source       string `json:"source"`
	HashTagId    int    `json:"hash_tag_id"`
	Sentiment    string `json:"sentiment"`
	FullText     string `gorm:"size:600" json:"full_text"`
	//// move tweet attributes
	NewSentiment string `json:"new_sentiment"`
	WhoMove      string `gorm:"index:who_move" json:"who_move"`
}

func (s *MigrationTables) TweetMigrate() {
	config.DB.AutoMigrate(&Tweet{})
}