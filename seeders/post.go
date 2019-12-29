package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"

	"syreclabs.com/go/faker"
)

func (s *Seeder) PostSeeder() {
	for i := 0; i < 10; i++ {
		newPost()
	}
}

func newPost() {
	data := models.Post{
		Title:      faker.RandomString(10),
		Body:       faker.RandomString(50),
		Image:      "bo47u2ro6higp2jdvuf0.jpeg",
		UserId:     uint(faker.RandomInt(1, 11)),
		CategoryId: uint(faker.RandomInt(1, 10)),
		Status:     int8(faker.RandomInt(1, 2)),
	}
	config.DB.Create(&data)
}
