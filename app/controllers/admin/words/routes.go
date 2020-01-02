package words

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	r.POST("words" , Store)
	r.PUT("words/:id" , Update)
	r.DELETE("words/:id" , Delete)
	return r
}
