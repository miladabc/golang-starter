package configtest

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"testing"
	"time"

	"github.com/miladabc/golang-starter/internal/config"
	"github.com/stretchr/testify/require"
)

func New(t *testing.T) *config.Config {
	t.Helper()

	cfg, err := config.New()
	require.NoError(t, err)

	cfg.DB.DBName += random(t)

	return cfg
}

func random(t *testing.T) string {
	t.Helper()

	b := make([]byte, 4)
	_, err := rand.Read(b)
	require.NoError(t, err)

	return "_" + hex.EncodeToString(b) + "_" + strconv.FormatInt(time.Now().Unix(), 10)
}
