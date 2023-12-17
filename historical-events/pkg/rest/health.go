package rest

import (
	"net/http"

	"github.com/aliaksandrrachko/historian/historical-events/api/rest/model"
	"github.com/aliaksandrrachko/historian/historical-events/pkg/health"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterHealthHandlers(router gin.IRouter, logger *logrus.Logger, health health.HealthCheck) {
	res := healthResource{logger, health}

	router.GET("/liveness", res.liveness)
	router.GET("/readiness", res.readiness)
}

type healthResource struct {
	logger      *logrus.Logger
	healthCheck health.HealthCheck
}

func (sR healthResource) liveness(c *gin.Context) {
	report := sR.healthCheck.Liveness()
	c.IndentedJSON(sR.mapToStatus(report.Status), sR.mapToApiModel(report))
}

func (sR healthResource) readiness(c *gin.Context) {
	report := sR.healthCheck.Readiness()
	c.IndentedJSON(sR.mapToStatus(report.Status), sR.mapToApiModel(report))
}

func (sR healthResource) mapToStatus(cs health.Status) (code int) {
	switch cs {
	case health.UNKNOWN:
		code = http.StatusInternalServerError
	case health.UP:
		code = http.StatusOK
	case health.DOWN:
		code = http.StatusServiceUnavailable
	case health.OUT_OF_SERVICE:
		code = http.StatusServiceUnavailable
	}
	return code
}

func (sR healthResource) mapToApiModel(hr health.HealthReport) model.HealthApiModel {
	apiComponents := make(map[string]model.ComponentApiModel)
	for _, component := range hr.Components {
		apiComponents[component.Description] = model.ComponentApiModel{Status: component.Status.String()}
	}

	healthApiModel := model.HealthApiModel{
		Status:     model.Status{Code: hr.Status.String()},
		Components: apiComponents,
	}

	return healthApiModel
}
