package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
)

func JWTAuthMiddleware(jwtAuth *jwtauth.JWTAuth) gin.HandlerFunc {
	verifier := jwtauth.Verifier(jwtAuth)

	return func(c *gin.Context) {

		handler := verifier(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
		}))

		handler.ServeHTTP(c.Writer, c.Request)

		_, claims, err := jwtauth.FromContext(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing token"})
			c.Abort()
			return
		}

		userID, exists := claims["sub"].(string)
		if !exists || userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in token"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}
