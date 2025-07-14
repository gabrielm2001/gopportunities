package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
	})
}

func ShowOpningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
	})
}

func UpdateOpningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
	})
}

func DeleteOpningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
	})
}

func ListOpningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
	})
}