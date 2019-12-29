package comments

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("comments", Index)
	r.POST("comments", Store)
	r.PUT("comments/:id", Update)
	r.DELETE("comments/:id", Delete)

	return r
}
