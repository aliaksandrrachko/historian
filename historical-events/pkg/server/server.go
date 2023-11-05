package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/aliaksandrrachko/historian/historical-events/pkg/config"
	"github.com/aliaksandrrachko/historian/historical-events/pkg/ginlogrus"
	statusApi "github.com/aliaksandrrachko/historian/historical-events/pkg/rest"
)

type HistoricalEventsServer interface {
	Start()
}

type historicalEventsServer struct {
	logger    *logrus.Logger
	config    config.Config
	ginEngine *gin.Engine
}

func NewServer(logger *logrus.Logger, config config.Config, ginEngine *gin.Engine) HistoricalEventsServer {
	return historicalEventsServer{logger: logger, config: config, ginEngine: ginEngine}
}

func (server historicalEventsServer) Start() {
	statusApi.RegisterHandlers(server.ginEngine.Group("/historical-events/api/v1"), server.logger)

	server.ginEngine.Run(fmt.Sprintf("%v:%v", server.config.Server.Host, server.config.Server.Port))
}

func NewEngine(logger *logrus.Logger, cfg config.Config) *gin.Engine {
	ginEngine := gin.New()
	ginEngine.Use(ginlogrus.Logger(logger))
	ginEngine.Use(gin.Recovery())
	ginEngine.SetTrustedProxies(cfg.Server.TrustedProxies)
	return ginEngine
}
