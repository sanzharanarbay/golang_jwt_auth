package middleware

import (
	"github.com/gin-gonic/gin"
	"jwt_auth_golang/application/handlers"
	"jwt_auth_golang/application/modules/auth"
	"jwt_auth_golang/application/services"
	"net/http"
)

var redisClient = services.InitRedis()
var rd = auth.NewAuth(redisClient)
var tk = auth.NewToken()
var service = handlers.NewProfile(rd, tk)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		metadata, err := service.Tk.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":       err,
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		_, err = service.Rd.FetchAuth(metadata.TokenUuid)
		if err != nil {
			c.JSON(http.StatusUnauthorized,
				gin.H{
					"error":       err,
					"message": "unauthorized",
				})
			c.Abort()
			return
		}


		c.Next()
	}
}


