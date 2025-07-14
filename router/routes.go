package router

import (
	"github.com/gabrielm2001/gopportunities.git/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handlers.CreateOpningHandler)
		v1.POST("/opening", handlers.CreateOpningHandler)
		v1.DELETE("/opening", handlers.DeleteOpningHandler)
		v1.PUT("/opening", handlers.UpdateOpningHandler)
		v1.GET("/openings", handlers.ListOpningHandler)
	}

}