package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware that handles errors attached to the context
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Execute the controller logic

		// Check if there are any errors attached to the context
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			// Use the status code if it was already set by the controller
			// Otherwise default to 500 Internal Server Error
			statusCode := c.Writer.Status()
			if statusCode == http.StatusOK {
				statusCode = http.StatusInternalServerError
			}
			
			c.JSON(statusCode, gin.H{"error": err.Error()})
		}
	}
}
