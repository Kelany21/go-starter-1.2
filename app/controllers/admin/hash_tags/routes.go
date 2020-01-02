package hash_tags

import "github.com/gin-gonic/gin"

/**
* all admin modules route will store here
*/
func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	r.GET("hash_tags" , Index)
	r.POST("hash_tags" , Store)
	r.PUT("hash_tags/:id" , Update)
	//r.POST("hash_tags/deactivate/:id" , Deactivate)
	r.POST("hash_tags/activate/:id" , Activate)
	r.POST("hash_tags/trash/:id" , Trash)
	r.GET("hash_tags/:id" , Show)
	r.DELETE("hash_tags/:id" , Delete)

	return r
}
