package routes

import (
	handlers "go-gin-boilerplate/server/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitLoginRoute(o, r, c *gin.RouterGroup) {
	o.POST("/login", handlers.LoginUser())
	o.POST("/register", handlers.RegisterUser())
}
