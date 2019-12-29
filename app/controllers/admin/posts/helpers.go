package posts

import (
	"fmt"
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/post"
	"golang-starter/config"
	"golang-starter/helpers"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func filter(g *gin.Context) []string {
	var filter []string
	if g.Query("title") != "" {
		filter = append(filter, `title like "%`+g.Query("title")+`%"`)
	}
	if g.Query("body") != "" {
		filter = append(filter, `body like "%`+g.Query("body")+`%"`)
	}
	if g.Query("user_id") != "" {
		filter = append(filter, "user_id = "+g.Query("user_id"))
	}
	if g.Query("category_id") != "" {
		filter = append(filter, "category_id = "+g.Query("category_id"))
	}
	if g.Query("status") != "" {
		filter = append(filter, "status = "+g.Query("status"))
	}

	return filter
}

func preload() []string {
	return []string{}
}

func validateRequest(g *gin.Context, action string) (bool, *models.Post) {
	var (
		category models.Category
		err      *govalidator.Validator
	)

	row := new(models.Post)

	if action == "store" {
		err = post.Store(g.Request, row)
	} else {
		err = post.Update(g.Request, row)
	}

	if helpers.ReturnNotValidRequest(err, g) {
		return false, row
	}

	if row.Image != "" {
		er := false
		er, row.Image = helpers.DecodeImage(g, row.Image)
		if !er {
			fmt.Println(row.Image)
			helpers.UploadError(g)
			return false, row
		}
	}

	if action == "store" {
		config.DB.Where("id = ?", row.CategoryId).First(&category)

		if category.ID == 0 {
			helpers.ReturnNotValidCategory(g)
			return false, row
		}
	}

	return true, row
}

func addUserToRow(g *gin.Context, row *models.Post) *models.Post {
	var user models.User

	adminToken := g.GetHeader("Authorization")
	config.DB.Where("token = ?", adminToken).First(&user)
	row.UserId = user.ID

	return row
}

func DeleteRelated(row models.Post) {
	var comments []models.Comment

	config.DB.Where("post_id = ?", row.ID).Delete(&comments)

}

func FindOrFail(id interface{}) (models.Post, bool) {
	var oldRow models.Post
	config.DB.Where("id = ? ", id).Find(&oldRow)
	if oldRow.ID != 0 {
		return oldRow, true
	}
	return oldRow, false
}

func updateColumns(row *models.Post, oldRow models.Post) models.Post {
	// if row.Image != "" {
	// }
	onlyAllowData := helpers.UpdateOnlyAllowColumns(row, models.PostFillAbleColumn())
	config.DB.Model(&oldRow).Updates(onlyAllowData)
	newData, _ := FindOrFail(oldRow.ID)
	return newData
}
