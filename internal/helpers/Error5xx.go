package helpers

import "github.com/gin-gonic/gin"

// Abort http req with json and 500 status code
func Error5xx(ctx *gin.Context, err error) {
	response := ResponseWrapper{
		Success: false,
		Error:   err,
		Data:    nil,
	}

	ctx.AbortWithStatusJSON(500, response)
}
