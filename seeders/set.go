package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"
	"syreclabs.com/go/faker"
)

func (s *Seeder) SetSeeder() {
	for i := 0 ; i < 70 ; i++ {
		newSet()
	}
}

/**
* fake data and create data base
 */
func newSet()  {
	data := models.Set{
		WordsCount: 0,
		UsedNumber: faker.RandomInt(1,100),
		UseCount:	faker.RandomInt(1,100),
		TrashedBy:	faker.Internet().UserName(),
		Name:		faker.Internet().UserName(),
		UserId:     faker.RandomInt(1,11),
		WhoAdd:   	faker.Internet().UserName(),
	}
	config.DB.Create(&data)
}
