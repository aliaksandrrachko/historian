package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/aliaksandrrachko/historian/historical-events/pkg/config"
	"github.com/aliaksandrrachko/historian/historical-events/pkg/ginlogrus"
	status "github.com/aliaksandrrachko/historian/historical-events/pkg/rest"
)

type HistoricalEventsServer interface {
	Start()
}

type historicalEventsServer struct {
	logger    *logrus.Logger
	config    config.Config
	ginEngine *gin.Engine
}

func New(logger *logrus.Logger, config config.Config, ginEngine *gin.Engine) HistoricalEventsServer {
	return historicalEventsServer{logger: logger, config: config, ginEngine: ginEngine}
}

func NewInstance() historicalEventsServer {
	logger := NewLogger()
	config := config.NewRandomConfig()
	ginEngine := NewEngine(logger, &config)
	return historicalEventsServer{logger: logger, config: config, ginEngine: ginEngine}
}

func (server historicalEventsServer) Start() {
	status.RegisterHandlers(server.ginEngine.Group("/historical-events/api/v1"), server.logger)

	server.ginEngine.Run(fmt.Sprintf("%v:%v", server.config.Server.Host, server.config.Server.Port))
}

func NewEngine(logger *logrus.Logger, cfg *config.Config) *gin.Engine {
	ginEngine := gin.New()
	ginEngine.Use(ginlogrus.Logger(logger))
	ginEngine.Use(gin.Recovery())
	ginEngine.SetTrustedProxies(cfg.Server.TrustedProxies)
	return ginEngine
}

func NewLogger() *logrus.Logger {
	// TODO: think change to singleton
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
