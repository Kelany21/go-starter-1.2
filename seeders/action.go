package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func (s *Seeder) ActionSeeder() {
	HashTagAction()
	UserAction()
	SetAction()
	PermissionAction()

}

/**
* fake data and create data base
 */
func HashTagAction() {
	actions := []models.Action{
		{
			Title: "all",
			Count:      70,
			ModuleName: "hash_tag",
		},

		{
			Title: "activated",
			Count:      70,
			ModuleName: "hash_tag",
		},

		{
			Title: "deactivated",
			Count:      0,
			ModuleName: "hash_tag",
		},

		{
			Title: "analyzed",
			Count:      0,
			ModuleName: "hash_tag",
		},


		{
			Title: "trashed",
			Count:      0,
			ModuleName: "hash_tag",
		},
	}
	for _, action := range actions {
		config.DB.Create(&action)
	}
}

func UserAction() {
	actions := []models.Action{
		{
			Title: "all",
			Count:      11,
			ModuleName: "user",
		},

		{
			Title: "activated",
			Count:      11,
			ModuleName: "user",
		},

		{
			Title: "deactivated",
			Count:      0,
			ModuleName: "user",
		},

		{
			Title: "trashed",
			Count:      0,
			ModuleName: "user",
		},

		{
			Title: "Blocked",
			Count:      0,
			ModuleName: "user",
		},
	}
	for _, action := range actions {
		config.DB.Create(&action)
	}
}

func SetAction() {
	actions := []models.Action{
		{
			Title:      "all",
			Count:      70,
			ModuleName: "set",
		},

		{
			Title:      "activated",
			Count:      70,
			ModuleName: "set",
		},

		{
			Title: "deactivated",
			Count:      0,
			ModuleName: "set",
		},

		{
			Title: "trashed",
			Count:      0,
			ModuleName: "set",
		},
	}
	for _, action := range actions {
		config.DB.Create(&action)
	}
}

func PermissionAction() {
	actions := []models.Action{
		{
			Title: "all",
			Count:      0,
			ModuleName: "permission",
		},

		{
			Title: "activated",
			Count:      0,
			ModuleName: "permission",
		},

		{
			Title: "deactivated",
			Count:      0,
			ModuleName: "permission",
		},

		{
			Title: "trashed",
			Count:      0,
			ModuleName: "permission",
		},
	}
	for _, action := range actions {
		config.DB.Create(&action)
	}
}

//TODO: the all action at any module should have title "all"
