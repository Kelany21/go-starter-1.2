package sets

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
)

func Index(g *gin.Context) {
	// array of rows and array of actions
	var (
		rows []models.Set
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
	config.DB.Where("module_name = ? " , "set").Find(&actions)
	// transform slice
	responseMap := make(map[string]interface{})
	responseMap["hash_tags"] = transformers.SetsResponse(rows)
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
	var count int
	config.DB.Model(models.Set{}).Where("name = ? ", row.Name).Count(&count)
	if count > 0 {
		helpers.ReturnDuplicateData(g, "name")
		return
	}
	row = addUserToRow(g, row)
	// create new row
	config.DB.Create(&row)
	// update all and activated actions counter
	UpdateAction()
	//now return row data after transformers
	helpers.OkResponse(g, helpers.DoneCreateItem(g), transformers.SetResponse(*row))
}

/***
* toggle the activation of row
 */
func Activate(g *gin.Context) {
	// get current row
	var row models.Set
	config.DB.Where("id = ?", g.Param("id")).Find(&row)
	// check if this row is deactivated
	if row.ActionId == 13 {
		// update activated and deactivated counter
		config.DB.Exec("UPDATE actions SET count = count + 1 WHERE id = ?", 12)
		config.DB.Exec("UPDATE actions SET count = count - 1 WHERE id = ?", 13)
		config.DB.Exec("UPDATE sets SET action_id = 12 WHERE id = ?", row.ID)
		//return activate response
		helpers.OkResponseWithOutData(g,helpers.DoneActivate(g))
		return
	}else
	// check if this row is activated
	if row.ActionId == 12 {
		// update activated and deactivated counter
		config.DB.Exec("UPDATE actions SET count = count + 1 WHERE id = ?", 13)
		config.DB.Exec("UPDATE actions SET count = count - 1 WHERE id = ?", 12)
		config.DB.Exec("UPDATE sets SET action_id = 13 WHERE id = ?", row.ID)
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
	var row models.Set
	config.DB.Where("id = ?", g.Param("id")).Find(&row)
	// check if this row is not trashed
	if row.ActionId != 14 {
		// update trashed and current action counter
		config.DB.Exec("UPDATE actions SET count = count + 1 WHERE id = ?", 14)
		config.DB.Exec("UPDATE actions SET count = count - 1 WHERE id = ?", row.ActionId)
		config.DB.Exec("UPDATE sets SET action_id = 14 WHERE id = ?", row.ID)
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
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneGetItem(g), transformers.SetResponse(row))
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
	if row.Name != oldRow.Name {
		var count int
		config.DB.Model(models.Set{}).Where("name = ? ", row.Name).Count(&count)
		if count > 0 {
			helpers.ReturnDuplicateData(g, "name")
			return
		}
	}

	/// update allow columns
	oldRow = updateColumns(row, oldRow)
	// now return row data after transformers
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.SetResponse(oldRow))
}

