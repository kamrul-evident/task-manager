package middleware

import (
	"net/http"
	"strings"
	"task-management/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware checks for a valid JWT in the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Expect "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte("h3Gz9aPqWmT1dU8kLzNr5FvR8yJx2Shq"), nil // Match jwtSecret from Login
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract claims and set in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", uint(claims["id"].(float64))) // Cast to uint
			c.Set("email", claims["email"].(string))
			c.Set("role", models.UserRole(claims["role"].(string)))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next() // Proceed to the next handler
	}
}


// package middleware

// import (
// 	"net/http"
// 	"os"

// 	"github.com/gin-gonic/gin"
// 	uuid "github.com/google/uuid"
// )

// // RequestIDMiddleware ...
// // Generate a unique ID and attach it to each request for future reference or use
// func RequestIDMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		uuid := uuid.New()
// 		c.Writer.Header().Set("X-Request-Id", uuid.String())
// 		c.Next()
// 	}
// }

// // TokenAuthMiddleware ...
// // JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
// func TokenAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authorization := c.Request.Header.Get("SECRET-KEY")
// 		if authorization == "" {
// 			//Token either expired or not valid
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token Verfication Failed"})
// 			return
// 		}
// 		if authorization != "6433220e-5f0b-4238-bb11-046f589e9149" {
// 			c.JSON(http.StatusUnauthorized, "Token Verfication Failed")
// 		}
// 		c.Next()
// 	}
// }

// func AuthAccessTokenMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		if c.Request.Header.Get("Authorization") != "" {
// 			client := &http.Client{}
// 			req, _ := http.NewRequest("GET", os.Getenv("AUTH_URL"), nil)
// 			req.Header.Set("Authorization", c.Request.Header.Get("Authorization"))
// 			resp, err := client.Do(req)
// 			if err != nil {
// 				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token Verfication Failed !"})
// 				return
// 			}
// 			defer resp.Body.Close()

// 			if resp.StatusCode != 200 {
// 				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token Verfication Failed !"})
// 				return
// 			}
// 			c.Next()
// 		}

// 		if c.Request.Header.Get("SECRET-KEY") != "" {
// 			if c.Request.Header.Get("SECRET-KEY") != "6433220e-5f0b-4238-bb11-046f589e9149" {
// 				c.AbortWithStatusJSON(http.StatusUnauthorized, "Token Verfication Failed")
// 				return
// 			}
// 			c.Next()
// 		}

// 	}
// }