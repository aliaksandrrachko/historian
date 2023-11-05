//go:build wireinject
// +build wireinject

package wiregen

import (
	"github.com/aliaksandrrachko/historian/historical-events/pkg/config"
	"github.com/aliaksandrrachko/historian/historical-events/pkg/ginlogrus"
	"github.com/aliaksandrrachko/historian/historical-events/pkg/server"
	"github.com/google/wire"
)

func InitServer() server.HistoricalEventsServer {
	panic(
		wire.Build(
			ginlogrus.NewLogger,
			config.NewRandomConfig,
			server.NewEngine,
			server.NewServer,
		),
	)
}
