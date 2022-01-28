package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth - generic health check
func GetHealth(c *gin.Context) error {
	response := HealthResponse{
		StatusCode: 200,
		Status:     "OK",
		Message:    "Health Check Successful",
	}

	c.JSON(http.StatusOK, response)
	return nil
}
