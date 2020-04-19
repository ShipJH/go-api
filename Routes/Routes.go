package Routes

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("member", Controllers.GetMembers)
		v1.POST("member", Controllers.CreateAMember)
		v1.GET("member/:id", Controllers.GetAMember)
		v1.PUT("member/:id", Controllers.UpdateAMember)
		v1.DELETE("member/:id", Controllers.DeleteAMember)
	}

	return r
}
