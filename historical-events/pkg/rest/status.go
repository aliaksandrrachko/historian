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
}

type statusResource struct {
	logger *logrus.Logger
}

func (sR statusResource) ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.StatusApiModel{Message: "pong", Timestamp: time.Now()})
}
