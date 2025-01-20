package container

import (
	"context"

	"github.com/jmoiron/sqlx"
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

func (c *Container) InitTodo() {
	todo.Init(c.HTTPRouter)
}

func (c *Container) Shutdown(ctx context.Context) {
	if c.HTTPServer != nil {
		c.HTTPServer.Shutdown(ctx)
	}

	if c.DB != nil {
		err := c.DB.Close()
		if err != nil {
			log.Error().Err(err).Msg("closing db connection")
		}
	}
}
