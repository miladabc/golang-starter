package containertest

import (
	"testing"

	"github.com/miladabc/golang-starter/internal/config/configtest"
	"github.com/miladabc/golang-starter/internal/container"
	"github.com/miladabc/golang-starter/internal/migrations"
	"github.com/miladabc/golang-starter/pkg/mysql/mysqltest"
	"github.com/stretchr/testify/require"
)

func New(t *testing.T) *container.Container {
	t.Helper()

	c := container.New()
	c.Config = configtest.New(t)
	c.DB = mysqltest.ConnectAndMigrate(t, c.Config.DB, migrations.FS, ".")

	err := c.Init()
	require.NoError(t, err)

	return c
}
