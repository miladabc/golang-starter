package app

import (
	"context"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/miladabc/golang-starter/internal/container/containertest"
	"github.com/stretchr/testify/require"
)

const (
	proto   = "http://"
	baseURL = "127.0.0.1:8080"
)

func TestE2E(t *testing.T) {
	t.Run("initialize app", func(t *testing.T) {
		c := containertest.New(t)

		t.Cleanup(func() {
			c.Shutdown(context.Background())
		})

		err := c.Init()
		require.NoError(t, err)

		go c.HTTPServer.Start()

		waitForServerStart(t)

		t.Run("check last todo: does not exist", func(t *testing.T) {
			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, proto+baseURL+"/api/v1/todo", nil)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			require.NoError(t, err)

			require.Contains(t, string(body), "Not Found")
		})
	})
}

func waitForServerStart(t *testing.T) {
	t.Helper()

	require.Eventually(t, func() bool {
		conn, err := net.Dial("tcp", baseURL)
		if err != nil || conn == nil {
			return false
		}

		err = conn.Close()

		return err == nil
	}, time.Second, time.Millisecond)
}
