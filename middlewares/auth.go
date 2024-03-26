package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"natthan.com/go-play/utils"
)

func Authenticate(context *gin.Context) {
	token := strings.Split(context.Request.Header.Get("Authorization"), " ")[1]
	fmt.Println(token)
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorize 1"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorize 2"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
