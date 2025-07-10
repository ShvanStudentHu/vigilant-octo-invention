package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


func Logger(logger *logrus.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
     
        start := time.Now()

       
        c.Next()

       
        duration := time.Since(start)

        logger.WithFields(logrus.Fields{
            "status":     c.Writer.Status(),
            "method":     c.Request.Method,
            "path":       c.Request.URL.Path,
            "duration":   duration,
            "client_ip":  c.ClientIP(),
            "user_agent": c.Request.UserAgent(),
        }).Info("HTTP request")
    }
}