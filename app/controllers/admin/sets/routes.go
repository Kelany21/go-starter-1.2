package sets

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	r.GET("sets" , Index)
	r.POST("sets" , Store)
	r.PUT("sets/:id" , Update)
	//r.POST("sets/deactivate/:id" , Deactivate)
	r.POST("sets/activate/:id" , Activate)
	r.POST("sets/trash/:id" , Trash)
	r.GET("sets/:id" , Show)
	r.DELETE("sets/:id" , Delete)

	return r
}
