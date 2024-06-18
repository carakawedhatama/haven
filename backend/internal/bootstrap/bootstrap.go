package bootstrap

import (
	"context"
	"fmt"
	"haven/pkg/config"
	"haven/pkg/validator"
	"os"
	"os/signal"
	"syscall"
	"time"

	"haven/internal/adapter/rest"

	"github.com/runsystemid/golog"
	"github.com/runsystemid/gontainer"
)

var appContainer = gontainer.New()

func Run(conf *config.Config) {
	appContainer.RegisterService("config", conf)

	// Initialize struct validator
	appContainer.RegisterService("validator", validator.NewGoValidator())

	bootstrapContext := context.Background()
	golog.Info(bootstrapContext, "Serving...")

	// Register adapter
	RegisterDatabase()
	RegisterCache()
	RegisterRest()
	RegisterToggleService()
	RegisterRepository()

	// Register application
	RegisterService()
	RegisterApi()

	// Startup the container
	if err := appContainer.Ready(); err != nil {
		golog.Panic(bootstrapContext, "Failed to populate service", err)
	}

	// Start server
	fiberApp := appContainer.GetServiceOrNil("fiber").(*rest.Fiber)
	errs := make(chan error, 2)
	go func() {
		golog.Info(bootstrapContext, fmt.Sprintf("Listening on port :%d", conf.Http.Port))
		errs <- fiberApp.Listen(fmt.Sprintf(":%d", conf.Http.Port))
	}()

	golog.Info(bootstrapContext, "Your app started")

	gracefulShutdown(bootstrapContext)
}

func gracefulShutdown(ctx context.Context) {
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	delay := 5 * time.Second

	golog.Info(ctx, fmt.Sprintf("Signal termination received. Waiting %v to shutdown.", delay))

	time.Sleep(delay)

	golog.Info(ctx, "Cleaning up resources...")

	// This will shuting down all the resources
	appContainer.Shutdown()

	golog.Info(ctx, "Bye")
}
