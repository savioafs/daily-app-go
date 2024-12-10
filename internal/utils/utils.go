package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(ctx *gin.Context) (string, bool) {
	userID, exists := ctx.Get("user_id")
	if !exists || userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user_id is required",
		})
		return "", false
	}

	id, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user_id",
		})
		return "", false
	}

	return id, true
}
