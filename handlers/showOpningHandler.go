package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
	})
}