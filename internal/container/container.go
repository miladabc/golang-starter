package container

import (
	"context"

	"github.com/miladabc/golang-starter/internal/config"
	"github.com/miladabc/golang-starter/internal/http"
	"github.com/miladabc/golang-starter/internal/log"
	"github.com/miladabc/golang-starter/internal/todo"
)

type Container struct {
	Config     *config.Config
	HTTPRouter *http.Router
	HTTPServer *http.Server
}

func New() *Container {
	return &Container{}
}

func (c *Container) Init() error {
	err := c.InitConfig()
	if err != nil {
		return err
	}

	err = c.InitLogger()
	if err != nil {
		return err
	}

	c.InitRouter()
	c.InitServer()

	c.InitTodo()

	return nil
}

func (c *Container) InitConfig() (err error) {
	c.Config, err = config.New()
	return
}

func (c *Container) InitLogger() error {
	return log.Init(c.Config.Log)
}

func (c *Container) InitRouter() {
	c.HTTPRouter = http.NewRouter(http.RouterConfig{Debug: c.Config.Debug})
}

func (c *Container) InitServer() {
	c.HTTPServer = http.NewServer(c.Config.Server, c.HTTPRouter.Handler())
}

func (c *Container) InitTodo() {
	todo.Init(c.HTTPRouter)
}

func (c *Container) Shutdown(ctx context.Context) {
	c.HTTPServer.Shutdown(ctx)
}
