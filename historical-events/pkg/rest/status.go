package status

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/aliaksandrrachko/historian/historical-events/api/rest/model"
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
	// TODO add versionService and real hash commit
	c.IndentedJSON(http.StatusOK, model.VersionApiModel{Version: "0.0.1-SNAPSHOT", GitCommit: "dummy-commit", GoVersion: "1.19"})
}
