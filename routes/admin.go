package routes

import (
	"golang-starter/app/controllers/admin/categories"
	"golang-starter/app/controllers/admin/comments"
	"golang-starter/app/controllers/admin/faqs"
	"golang-starter/app/controllers/admin/posts"
	"golang-starter/app/controllers/admin/settings"
	"golang-starter/app/controllers/admin/translations"
	"golang-starter/app/controllers/admin/users"
	"golang-starter/app/controllers/visitor/pages"

	"github.com/gin-gonic/gin"
)

/***
* any route here will add after /admin
* admin only  will have access this routes
 */
func Admin(r *gin.RouterGroup) *gin.RouterGroup {
	categories.Routes(r)
	translations.Routes(r)
	settings.Routes(r)
	users.Routes(r)
	pages.Routes(r)
	faqs.Routes(r)
	posts.Routes(r)
	comments.Routes(r)

	return r
}
