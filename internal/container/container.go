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
	c.Config, err = config.New()
	return
}

func (c *Container) InitLogger() error {
	return ilog.Init(c.Config.Log)
}

func (c *Container) InitMySQL() (err error) {
	c.DB, err = mysql.Connect(c.Config.DB)
	return
}

func (c *Container) InitRouter() {
	c.HTTPRouter = http.NewRouter(http.RouterConfig{Debug: c.Config.Debug})
}

func (c *Container) InitServer() {
	c.HTTPServer = http.NewServer(c.Config.Server, c.HTTPRouter.Handler())
}

func (c *Container) InitSwagger() {
	docs.RegisterRoutes(c.HTTPRouter.Root)
}

func (c *Container) InitTodo() {
	c.Todo = todo.New(c.DB)
	c.Todo.RegisterRoutes(c.HTTPRouter.V1)
}

func (c *Container) Shutdown(ctx context.Context) {
	if c.HTTPServer != nil {
		c.HTTPServer.Shutdown(ctx)
	}

	if c.Todo != nil {
		c.Todo.Shutdown()
	}

	if c.DB != nil {
		err := c.DB.Close()
		if err != nil {
			log.Error().Err(err).Msg("closing db connection")
		}
	}
}
