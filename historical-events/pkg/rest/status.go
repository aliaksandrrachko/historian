package status

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/aliaksandrrachko/historian/historical-events/api/rest/model"
	"github.com/aliaksandrrachko/historian/historical-events/internal/build"
)

func RegisterHandlers(router gin.IRouter, logger *logrus.Logger) {
	res := statusResource{logger}

	router.GET("/status/ping", res.ping)
	router.GET("/status/version", res.version)
}

type statusResource struct {
	logger *logrus.Logger
}

func (sR statusResource) ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.StatusApiModel{Message: "pong", Timestamp: time.Now()})
}

func (sR statusResource) version(c *gin.Context) {
	buildInfo := build.Get()
	c.IndentedJSON(http.StatusOK, model.VersionApiModel{
		Version:   buildInfo.Version,
		GitCommit: buildInfo.GitCommit,
		GoVersion: buildInfo.GoVersion,
	})
}
