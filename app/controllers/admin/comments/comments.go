package comments

import (
	"golang-starter/app/models"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"

	"github.com/gin-gonic/gin"
)

func Index(g *gin.Context) {
	// array of rows
	var rows []models.Comment
	// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      config.DB,
		Page:    helpers.Page(g),
		Limit:   helpers.Limit(g),
		OrderBy: helpers.Order("id desc"),
		Filters: filter(g),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)
	// transform slice
	paginator.Records = transformers.CommentsResponse(rows)
	// return response
	helpers.OkResponseWithPaging(g, helpers.DoneGetAllItems(g), paginator)
}

func Store(g *gin.Context) {
	// check if request valid
	valid, row := validateRequest(g, "store")
	if !valid {
		return
	}
	row = addUserToRow(g, row)
	// create new row
	config.DB.Create(&row)
	//now return row data after transformers
	helpers.OkResponse(g, helpers.DoneCreateItem(g), transformers.CommentResponse(*row))
}

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

func Update(g *gin.Context) {
	// check if request valid
	valid, row := validateRequest(g, "update")
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
	helpers.OkResponse(g, helpers.DoneUpdate(g), transformers.CommentResponse(oldRow))
}
