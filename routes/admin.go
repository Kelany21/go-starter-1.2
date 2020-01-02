package routes

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/controllers/admin/Pages"
	"golang-starter/app/controllers/admin/categories"
	"golang-starter/app/controllers/admin/faqs"
	"golang-starter/app/controllers/admin/sets"
	"golang-starter/app/controllers/admin/settings"
	"golang-starter/app/controllers/admin/translations"
	"golang-starter/app/controllers/admin/users"
	"golang-starter/app/controllers/admin/hash_tags"
	"golang-starter/app/controllers/admin/words"
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
	hash_tags.Routes(r)
	sets.Routes(r)
	words.Routes(r)

	return r
}
