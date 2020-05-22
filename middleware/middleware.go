package middleware

import (
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := oauth.AuthenticateRequest(c.Request); err != nil {
			c.JSON(err.Status(), err)
			c.Abort()
			return
		}

		if userID := oauth.GetCallerID(c.Request); userID == 0 {
			restErr := rest_errors.NewUnauthorizedError("invalid credentials")
			c.JSON(restErr.Status(), restErr)
			c.Abort()
			return
		}

		c.Next()
	}
}
