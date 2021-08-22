package middleware

import (
	"fmt"

	"go-gin-boilerplate/server/api/models"
	routes "go-gin-boilerplate/server/api/routes"
	services "go-gin-boilerplate/server/api/services"
	jwthelper "go-gin-boilerplate/server/utils/jwthelper"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init -Init
func InitMiddleware(g *gin.Engine) {
	g.Use(cors.Default()) // CORS request
	o := g.Group("/o")
	o.Use(OpenRequestMiddleware())
	r := g.Group("/r")
	r.Use(RestrictedRequestMiddleware())
	// r.Use(jwt.Auth(models.JWTKey))
	c := r.Group("/c")
	c.Use(RoleBasedRequestMiddleware())
	routes.InitLoginRoute(o, r, c)
}

func OpenRequestMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("OpenRequestMiddleware called")
	}
}

// Need to check JWT token here

func RestrictedRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.Login{}
		// c.Bind(&user)
		token := c.GetHeader("Authorization")

		login, err := jwthelper.GetLoginFromToken(c)
		if err != nil {
			fmt.Println("Token not available", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}
		if strings.Trim(token, "") == "" {
			fmt.Println("Token not available")
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}
		_, isValid, usererr := services.ValidateUser(login)
		if usererr != nil || !isValid {
			fmt.Println("Failed to validate user")
			c.AbortWithStatusJSON(401, gin.H{"error": "Failed to validate user"})
		} else {
			user.UserName = login.UserName
			user.Password = login.Password
			// c.JSON(http.StatusOK, "User authentication successfull")
			c.Next()
		}

		fmt.Println("RestrictedRequestMiddleware called")
	}
}

// Need to check JWT token here with group validation
func RoleBasedRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("RoleBasedRequestMiddleware called")
	}
}
