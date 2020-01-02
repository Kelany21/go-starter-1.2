package hash_tags

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-starter/app/controllers/admin/words"
	"golang-starter/app/models"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
)

/***
* get all rows with pagination
 */
func Index(g *gin.Context) {
	// array of rows and array of actions
	var (
		rows []models.HashTag
		actions []models.Action
	)
	fmt.Println("here",g.Query("order"))
	// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      config.DB,
		Page:    helpers.Page(g),
		Limit:   helpers.Limit(g),
		OrderBy: helpers.Order(g,"id desc"),
		Filters: filter(g),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)
	// get model actions
	config.DB.Where("module_name = ? " , "hash_tag").Find(&actions)
	// transform slice
	responseMap := make(map[string]interface{})
	responseMap["hash_tags"] = transformers.HashTagsResponse(rows)
	responseMap["actions"] = transformers.ActionsResponse(actions)
	paginator.Records = responseMap

	// return response
	helpers.OkResponseWithPaging(g, helpers.DoneGetAllItems(g), paginator)
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
	row = addUserToRow(g, row)
	// create new row
	config.DB.Create(&row)
	// update all and activated actions counter
	UpdateAction()
	//now return row data after transformers
	helpers.OkResponse(g, helpers.DoneCreateItem(g), transformers.HashTagResponse(*row))
}

/***
* toggle the activation of row
 */
func Activate(g *gin.Context) {
	// get current row
	var row models.HashTag
	config.DB.Where("id = ?", g.Param("id")).Find(&row)
	// check if this row is deactivated
	if row.ActionId == 3 {
		// update activated and deactivated counter
		config.DB.Exec("UPDATE actions SET count = count + 1 WHERE id = ?", 2)
		config.DB.Exec("UPDATE actions SET count = count - 1 WHERE id = ?", 3)
		config.DB.Exec("UPDATE hash_tags SET action_id = 2 WHERE id = ?", row.ID)
		//return activate response
		helpers.OkResponseWithOutData(g,helpers.DoneActivate(g))
		return
	}else
	// check if this row is activated
	if row.ActionId == 2 {
		// update activated and deactivated counter
		config.DB.Exec("UPDATE actions SET count = count + 1 WHERE id = ?", 3)
		config.DB.Exec("UPDATE actions SET count = count - 1 WHERE id = ?", 2)
		config.DB.Exec("UPDATE hash_tags SET action_id = 3 WHERE id = ?", row.ID)
		// return deactivate response
		helpers.OkResponseWithOutData(g,helpers.DoneDeactivate(g))
		return
	}
	// return something went wrong response
	helpers.OkResponseWithOutData(g,helpers.Wrong(g))
	return
}

/***
* Trashing row
 */
func Trash(g *gin.Context) {
	// get current row
	var row models.HashTag
	config.DB.Where("id = ?", g.Param("id")).Find(&row)
	// check if this row is not trashed
	if row.ActionId != 5 {
		// update trashed and current action counter
		config.DB.Exec("UPDATE actions SET count = count + 1 WHERE id = ?", 5)
		config.DB.Exec("UPDATE actions SET count = count - 1 WHERE id = ?", row.ActionId)
		config.DB.Exec("UPDATE hash_tags SET action_id = 5 WHERE id = ?", row.ID)
	}
	// return trashed response
	helpers.OkResponseWithOutData(g,helpers.DoneTrash(g))
}

/***
* return row with id
 */
func Show(g *gin.Context) {
	// find this row or return 404
	row, find := FindOrFail(g.Param("id"))
	if !find {
		helpers.ReturnNotFound(g, helpers.ItemNotFound(g))
		return
	}
	if g.Query("with_words") != ""{
		data := make(map[string]interface{})
		data["words"] = transformers.WordsResponse(words.Index(g))
		data["records"] = transformers.HashTagResponse(row)
		// now return row data after transformers
		helpers.OkResponse(g, helpers.DoneGetItem(g), data)
		return
	}
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneGetItem(g), transformers.HashTagResponse(row))
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
	// update all and current actions counter
	action := row.ActionId
	config.DB.Unscoped().Delete(&row)
	UpdateActionAfterDelete(action)
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
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.HashTagResponse(oldRow))
}
