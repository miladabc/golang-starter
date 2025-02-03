package container

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/miladabc/golang-starter/docs"
	"github.com/miladabc/golang-starter/internal/config"
	"github.com/miladabc/golang-starter/internal/http"
	ilog "github.com/miladabc/golang-starter/internal/log"
	"github.com/miladabc/golang-starter/internal/todo"
	"github.com/miladabc/golang-starter/pkg/mysql"
	"github.com/rs/zerolog/log"
)

type Container struct {
	Config     *config.Config
	DB         *sqlx.DB
	HTTPRouter *http.Router
	HTTPServer *http.Server
	Todo       *todo.App
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

	err = c.InitMySQL()
	if err != nil {
		return err
	}

	c.InitRouter()
	c.InitServer()
	c.InitSwagger()

	c.InitTodo()

	return nil
}

func (c *Container) InitConfig() (err error) {
	if notNil(c.Config) {
		return
	}

	c.Config, err = config.New()

	return
}

func (c *Container) InitLogger() error {
	return ilog.Init(c.Config.Log)
}

func (c *Container) InitMySQL() (err error) {
	if notNil(c.DB) {
		return
	}

	c.DB, err = mysql.Connect(c.Config.DB)

	return
}

func (c *Container) InitRouter() {
	if notNil(c.HTTPRouter) {
		return
	}

	c.HTTPRouter = http.NewRouter(http.RouterConfig{Debug: c.Config.Debug})
}

func (c *Container) InitServer() {
	if notNil(c.HTTPServer) {
		return
	}

	c.HTTPServer = http.NewServer(c.Config.Server, c.HTTPRouter.Handler())
}

func (c *Container) InitSwagger() {
	docs.RegisterRoutes(c.HTTPRouter.Root)
}

func (c *Container) InitTodo() {
	if notNil(c.Todo) {
		return
	}

	c.Todo = todo.New(c.DB)
	c.Todo.RegisterRoutes(c.HTTPRouter.V1)
}

func (c *Container) Shutdown(ctx context.Context) {
	if notNil(c.HTTPServer) {
		c.HTTPServer.Shutdown(ctx)
	}

	if notNil(c.Todo) {
		c.Todo.Shutdown()
	}

	if notNil(c.DB) {
		err := c.DB.Close()
		if err != nil {
			log.Error().Err(err).Msg("closing db connection")
		}
	}
}

func notNil[T any](p *T) bool {
	return p != nil
}
