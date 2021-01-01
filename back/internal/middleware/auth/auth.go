package auth

import (
"github.com/gin-gonic/gin"
)

const authHeaderKey = "Authorization"
const usernameKey = "username"
type contextKey struct {
	name string
}


func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authHeaderKey)
		if header==""{
			c.Next()
			return
		}
		// Set username in the ggcontext
		//username, err := jwt.ParseToken(header)
		//if err!=nil {
		//	c.Status(http.StatusUnauthorized)
		//	return
		//}
		//c.Set("username", username)
		c.Set(usernameKey,header)
		c.Next()
	}
}


// ForContext finds the user from the ggcontext. REQUIRES Middleware to have run.
func ForContext(ctx *gin.Context) string {
	return ctx.GetString(usernameKey)
}