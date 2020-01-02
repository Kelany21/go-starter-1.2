package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"
	"syreclabs.com/go/faker"
)

func (s *Seeder) WordSeeder() {
	for i := 0 ; i < 7000 ; i++ {
		newSetWord()
		newHashTagWord()
	}
}

/**
* fake data and create data base
*/
func newSetWord()  {
	var Type = []string{"negative", "neutral", "positive"}
	data := models.Word{
		Word:		faker.Internet().UserName(),
		SourceType: "sets",
		SourceId: 	faker.RandomInt(1,70),
		Type:		Type[faker.RandomInt(0,2)],
	}
	config.DB.Exec("UPDATE sets SET words_count = words_count + 1 WHERE id = ?", data.SourceId)
	config.DB.Create(&data)
}

func newHashTagWord()  {
	var Type = []string{"negative", "neutral", "positive"}
	data := models.Word{
		Word:		faker.Internet().UserName(),
		SourceType: "hash_tags",
		SourceId: 	faker.RandomInt(1,70),
		Type:		Type[faker.RandomInt(0,2)],
	}
	config.DB.Exec("UPDATE hash_tags SET words_count = words_count + 1 WHERE id = ?", data.SourceId)
	config.DB.Create(&data)
}
