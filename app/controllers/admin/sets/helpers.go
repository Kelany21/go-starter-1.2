package sets

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/set"
	"golang-starter/config"
	"golang-starter/helpers"
)

func filter(g *gin.Context) []string {
	var filter []string
	if g.Query("action_id") != "" && g.Query("action_id") != "all" {
		filter = append(filter, "action_id = "+g.Query("action_id"))
	}
	if g.Query("name") != "" {
		filter = append(filter, `name like "%`+g.Query("name") + `%"`)
	}
	if g.Query("from") != "" && g.Query("to") != "" {
		filter = append(filter, `created_at BETWEEN "` + g.Query("from") + `" AND "` + g.Query("to") + `"`)
	}else if g.Query("from") != "" {
		filter = append(filter, `created_at  >= "`+ g.Query("from") + `"`)
	}else if g.Query("to") != "" {
		filter = append(filter, `created_at  <= "` + g.Query("to") + `"`)
	}
	if g.Query("who_add") != "" {
		filter = append(filter, `who_add like "%`+g.Query("who_add") + `%"`)
	}
	return filter
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
func validateRequest(g *gin.Context) (bool, *models.Set) {
	// init struct to validate request
	row := new(models.Set)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := set.StoreUpdate(g.Request, row)
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
* here we add auth user to row
 */
func addUserToRow(g *gin.Context, row *models.Set) *models.Set {
	// get auth user
	var user models.User
	adminToken := g.GetHeader("Authorization")
	config.DB.Where("token = ?", adminToken).First(&user)
	// set auth user id and name to row
	row.UserId = int(user.ID)
	row.WhoAdd = user.Name
	// return the row after adding the auth user
	return row
}

/**
* findOrFail Data
 */
func FindOrFail(id interface{}) (models.Set, bool) {
	var oldRow models.Set
	config.DB.Where("id = ?", id).Find(&oldRow)
	if oldRow.ID != 0 {
		return oldRow, true
	}
	return oldRow, false
}

/**
* increase all and activated actions counter
 */
func UpdateAction() {
	config.DB.Exec("UPDATE actions SET count = count + 1 WHERE id IN (?)", []string{"11","12"})
}

/**
* decrease all and current actions counter
 */
func UpdateActionAfterDelete(actionId int) {
	config.DB.Exec("UPDATE actions SET count = count - 1 WHERE id IN (?)", []int{11, actionId})
}

/**
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
 */
func updateColumns(row *models.Set, oldRow models.Set) models.Set {
	onlyAllowData := helpers.UpdateOnlyAllowColumns(row, models.SetFillAbleColumn())
	/// disable auto association (we will update only what we need)
	config.DB.Model(&oldRow).Updates(onlyAllowData)
	/// return data with preload relations
	newData, _ := FindOrFail(oldRow.ID)

	return newData
}
