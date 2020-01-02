package words

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/word"
	"golang-starter/config"
	"golang-starter/helpers"
)

func filter(g *gin.Context) []string {
	return []string {}
}

/**
* preload module with some preload conditions
 */
func preload() []string {
	return []string{}
}

/**
* here we will check if request valid or not
 */
func validateRequest(g *gin.Context) (bool, *models.Word) {
	// init struct to validate request
	row := new(models.Word)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := word.StoreUpdate(g.Request, row)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return false, row
	}
	return true, row
}

/**
* findOrFail Data
 */
func FindOrFail(id interface{}) (models.Word, bool) {
	var oldRow models.Word
	config.DB.Where("id = ?", id).Find(&oldRow)
	if oldRow.ID != 0 {
		return oldRow, true
	}
	return oldRow, false
}

/**
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
 */
func updateColumns(row *models.Word, oldRow models.Word) models.Word {
	onlyAllowData := helpers.UpdateOnlyAllowColumns(row, models.WordFillAbleColumn())
	/// disable auto association (we will update only what we need)
	config.DB.Model(&oldRow).Updates(onlyAllowData)
	/// return data with preload relations
	newData, _ := FindOrFail(oldRow.ID)

	return newData
}

