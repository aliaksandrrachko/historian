package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/aliaksandrrachko/historian/historical-events/pkg/config"
	"github.com/aliaksandrrachko/historian/historical-events/pkg/ginlogrus"
	"github.com/aliaksandrrachko/historian/historical-events/pkg/health"
	apiResource "github.com/aliaksandrrachko/historian/historical-events/pkg/rest"
)

type HistoricalEventsServer interface {
	Start()
}

type historicalEventsServer struct {
	logger      *logrus.Logger
	config      config.Config
	ginEngine   *gin.Engine
	healthCheck health.HealthCheck
}

func NewServer(
	logger *logrus.Logger,
	config config.Config,
	ginEngine *gin.Engine,
	healthCheck health.HealthCheck,
) HistoricalEventsServer {
	return historicalEventsServer{logger: logger, config: config, ginEngine: ginEngine, healthCheck: healthCheck}
}

func (server historicalEventsServer) Start() {
	apiResource.RegisterStatusHandlers(server.ginEngine.Group("/historical-events/api/v1/status"), server.logger)
	apiResource.RegisterHealthHandlers(server.ginEngine.Group("/historical-events/api/v1/health"), server.logger, server.healthCheck)

	if err := server.ginEngine.Run(fmt.Sprintf("%v:%v", "", server.config.Server.Port)); err != nil {
		panic(err)
	}
}

func NewEngine(logger *logrus.Logger, cfg config.Config) *gin.Engine {
	ginEngine := gin.New()
	ginEngine.Use(ginlogrus.Logger(logger))
	ginEngine.Use(gin.Recovery())

	if err := ginEngine.SetTrustedProxies(cfg.Server.TrustedProxies); err != nil {
		panic(err)
	}

	return ginEngine
}
