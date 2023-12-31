package ginlogrus

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(logger logrus.FieldLogger, notLogged ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		requestQueryParams := c.Request.URL.Query()
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := logger.WithFields(logrus.Fields{
			"hostname":           hostname,
			"statusCode":         statusCode,
			"latency":            latency, // time to process
			"clientIP":           clientIP,
			"method":             c.Request.Method,
			"path":               path,
			"requestQueryParams": requestQueryParams,
			"referer":            referer,
			"dataLength":         dataLength,
			"userAgent":          clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf(
				"%s - %s \"%s %s\" (%dms)",
				clientIP,
				hostname,
				c.Request.Method,
				path,
				latency,
			)

			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}

func NewLogger() *logrus.Logger {
	return &logrus.Logger{
		Out: os.Stderr,
		// %d{HH:mm:ss.SSS} %-5level {%thread} [%logger{20}] : %msg%n
		Formatter: &logrus.TextFormatter{
			ForceColors:   true,
			DisableColors: false,
			FullTimestamp: true,
		},
		Hooks: make(logrus.LevelHooks),
		Level: logrus.DebugLevel,
	}
}
