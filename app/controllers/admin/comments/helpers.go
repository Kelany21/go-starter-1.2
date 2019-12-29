package comments

import (
	"golang-starter/app/models"
	"golang-starter/app/requests/admin/comment"
	"golang-starter/config"
	"golang-starter/helpers"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func filter(g *gin.Context) []string {
	var filter []string
	if g.Query("status") != "" {
		filter = append(filter, "status = "+g.Query("status"))
	}
	if g.Query("comment") != "" {
		filter = append(filter, `comment like "%`+g.Query("comment")+`%"`)
	}
	return filter
}

func preload() []string {
	return []string{}
}

func validateRequest(g *gin.Context, action string) (bool, *models.Comment) {
	var (
		post models.Post
		err  *govalidator.Validator
	)

	row := new(models.Comment)

	if action == "store" {
		err = comment.Store(g.Request, row)
	} else {
		err = comment.Update(g.Request, row)
	}

	if helpers.ReturnNotValidRequest(err, g) {
		return false, row
	}

	if action == "store" {
		config.DB.Where("id = ?", row.PostId).First(&post)

		if post.ID == 0 {
			helpers.ReturnNotValidPost(g)
			return false, row
		}
	}
	return true, row
}

func addUserToRow(g *gin.Context, row *models.Comment) *models.Comment {
	var user models.User

	adminToken := g.GetHeader("Authorization")
	config.DB.Where("token = ?", adminToken).First(&user)
	row.UserId = user.ID

	return row
}

func FindOrFail(id interface{}) (models.Comment, bool) {
	var oldRow models.Comment
	config.DB.Where("id = ?", id).Find(&oldRow)
	if oldRow.ID != 0 {
		return oldRow, true
	}
	return oldRow, false
}

func updateColumns(row *models.Comment, oldRow models.Comment) models.Comment {
	onlyAllowData := helpers.UpdateOnlyAllowColumns(row, models.CommentFillAbleColumn())
	config.DB.Model(&oldRow).Updates(onlyAllowData)
	newData, _ := FindOrFail(oldRow.ID)
	return newData
}
