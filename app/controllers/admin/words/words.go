package words

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
)

func Index(g *gin.Context) []models.Word {
	// array of words of source
	var rows []models.Word
	config.DB.Where("source_id = ?",g.Param("id")).Find(&rows)
	return rows
}

/**
* store new user
 */
func Store(g *gin.Context) {
	// check if request valid
	valid, row := validateRequest(g)
	if !valid {
		return
	}
	// create new row
	config.DB.Create(&row)
	if row.SourceType == "hash_tags" {
		config.DB.Exec("UPDATE SET words_count = words_count + 1 WHERE id = ?", row.SourceId)
	}else {
		config.DB.Exec("UPDATE sets SET words_count = words_count + 1 WHERE id = ?", row.SourceId)
	}
	//now return row data after transformers
	helpers.OkResponse(g, helpers.DoneCreateItem(g), transformers.WordResponse(*row))
}

/***
* delete row with id
 */
func Delete(g *gin.Context) {
	// find this row or return 404
	row, find := FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	config.DB.Unscoped().Delete(&row)
	// now return row data after transformers
	helpers.OkResponseWithOutData(g, helpers.DoneDelete(g))
}

/**
* update user
 */
func Update(g *gin.Context) {
	// check if request valid
	valid, row := validateRequest(g)
	if !valid {
		return
	}
	// find this row or return 404
	oldRow, find := FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	/// update allow columns
	oldRow = updateColumns(row, oldRow)
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.WordResponse(oldRow))
}
