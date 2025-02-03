package mysqltest

import (
	"fmt"
	"io/fs"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/miladabc/golang-starter/pkg/mysql"
	"github.com/stretchr/testify/require"
)

func ConnectAndMigrate(t *testing.T, cfg mysql.Config, migrateFS fs.FS, migratePath string) *sqlx.DB {
	t.Helper()

	grantPrivileges(t, cfg)

	err := mysql.CreateDB(cfg)
	require.NoError(t, err)

	t.Cleanup(func() {
		err := mysql.DropDB(cfg)
		if err != nil {
			t.Logf("dropping db: %s", err)
		}
	})

	db, err := mysql.Connect(cfg)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = db.Close()
		if err != nil {
			t.Logf("closing db connection: %s", err)
		}
	})

	m, err := mysql.NewMigration(mysql.MigrationConfig{
		FS:     migrateFS,
		Path:   migratePath,
		DB:     db.DB,
		DBName: cfg.DBName,
	})
	require.NoError(t, err)

	err = m.Up()
	require.NoError(t, err)

	return db
}

func grantPrivileges(t *testing.T, cfg mysql.Config) {
	t.Helper()

	rootCfg := cfg
	rootCfg.User = "root"
	rootCfg.Password = "root"
	rootCfg.DBName = ""

	db, err := mysql.Connect(rootCfg)
	require.NoError(t, err)
	defer db.Close()

	query := fmt.Sprintf("GRANT ALL PRIVILEGES ON *.* TO '%s';", cfg.User)
	_, err = db.Exec(query)
	require.NoError(t, err)
}
