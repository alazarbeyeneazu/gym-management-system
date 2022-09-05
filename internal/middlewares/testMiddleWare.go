package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"success": "aborted",
		})
		log.Println("success")
		return
	}
}
