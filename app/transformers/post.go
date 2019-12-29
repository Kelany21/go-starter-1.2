package transformers

import (
	"golang-starter/app/models"
	"golang-starter/config"
)

func PostResponse(post models.Post) map[string]interface{} {
	var (
		u        = make(map[string]interface{})
		category models.Category
		user     models.User
	)

	config.DB.Where("id = ?", post.CategoryId).First(&category)
	config.DB.Where("id = ?", post.UserId).First(&user)

	u["id"] = post.ID
	u["title"] = post.Title
	u["body"] = post.Body
	u["image"] = post.Image
	u["user"] = user.Name
	u["category"] = category.Name
	u["status"] = post.Status

	return u
}

/**
* stander the Multi users response
 */
func PostsResponse(posts []models.Post) []map[string]interface{} {
	var u = make([]map[string]interface{}, 0)
	for _, post := range posts {
		u = append(u, PostResponse(post))
	}
	return u

}
