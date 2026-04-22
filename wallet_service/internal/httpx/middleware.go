package httpx

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ServerID() gin.HandlerFunc {
	hostname, _ := os.Hostname()
	return func(c *gin.Context) {
		c.Header("X-Server-ID", hostname)
		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}
		err := c.Errors.Last().Err

		var he *HTTPError
		if errors.As(err, &he) {
			c.JSON(he.Status, gin.H{"error": he.Message})
			return
		}

		log.Printf("internal error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
}
