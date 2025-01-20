package app

import (
	"context"

	"github.com/miladabc/golang-starter/internal/container"
)

func StartHTTPServer() error {
	ctx := context.Background()
	c := container.New()
	defer c.Shutdown(ctx)

	err := c.Init()
	if err != nil {
		return err
	}

	go c.HTTPServer.Start()

	<-handleInterrupts()

	return nil
}
