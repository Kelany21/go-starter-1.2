package posts

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("posts", Index)
	r.POST("posts", Store)
	r.PUT("posts/:id", Update)
	r.GET("posts/:id", Show)
	r.DELETE("posts/:id", Delete)

	return r
}
