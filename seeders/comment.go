package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"

	"syreclabs.com/go/faker"
)

func (s *Seeder) CommentSeeder() {
	for i := 0; i < 10; i++ {
		newComment()
	}
}

/**
* fake data and create data base
 */
func newComment() {
	data := models.Comment{
		Comment: faker.RandomString(50),
		UserId:  uint(faker.RandomInt(1, 11)),
		PostId:  uint(faker.RandomInt(1, 10)),
		Status:  int8(faker.RandomInt(1, 2)),
	}
	config.DB.Create(&data)
}
