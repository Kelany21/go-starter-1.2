package transformers

import (
	"golang-starter/app/models"
	"golang-starter/config"
)

func CommentResponse(comment models.Comment) map[string]interface{} {
	var (
		u    = make(map[string]interface{})
		post models.Post
		user models.User
	)

	config.DB.Where("id = ?", comment.PostId).First(&post)
	config.DB.Where("id = ?", comment.UserId).First(&user)

	u["id"] = comment.ID
	u["user"] = user.Name
	u["post"] = post.Title
	u["comment"] = comment.Comment
	u["status"] = comment.Status

	return u
}

func CommentsResponse(comments []models.Comment) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, comment := range comments {
		u = append(u, CommentResponse(comment))
	}

	return u
}
