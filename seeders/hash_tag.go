package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"
	"syreclabs.com/go/faker"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func (s *Seeder) HashTagSeeder() {
	for i := 0 ; i < 70 ; i++ {
		newHashTag()
	}
}

/**
* fake data and create data base
 */
func newHashTag()  {
	data := models.HashTag{
		HashTag: 			faker.Internet().UserName(),
		GoodCount: 			faker.RandomInt(1,100),
		BadCount: 		  	faker.RandomInt(1,100),
		TweetCount:   		faker.RandomInt(1,100),
		WordsCount:   		0,
		NeutralCount: 		faker.RandomInt(1,100),
		SetsCount:  		faker.RandomInt(1,70),
		BadWordCount:  		faker.RandomInt(1,100),
		GoodWordCount:  	faker.RandomInt(1,100),
		NeutralWordCount:   faker.RandomInt(1,100),
		UserId:     		faker.RandomInt(1,11),
		WhoAdd:   			faker.Internet().UserName(),
		CountMoveGood:   	faker.RandomInt(1,100),
		CountMoveBad:    	faker.RandomInt(1,100),
		CountMoveNeutral: 	faker.RandomInt(1,100),
		Stream: 			faker.RandomInt(1,10),
	}
	config.DB.Create(&data)
}
