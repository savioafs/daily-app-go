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

		if claims["sub"] == nil || claims["sub"] == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token payload"})
			c.Abort()
			return
		}

		c.Next()
	}
}
