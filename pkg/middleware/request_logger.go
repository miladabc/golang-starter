package middleware

import (
	"strconv"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func AccessRequestLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogLatency:       true,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogURIPath:       false,
		LogRoutePath:     true,
		LogRequestID:     true,
		LogReferer:       true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogError:         true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Int64("latency", int64(v.Latency)).
				Str("latency_human", v.Latency.String()).
				Str("protocol", v.Protocol).
				Str("remote_ip", v.RemoteIP).
				Str("host", v.Host).
				Str("method", v.Method).
				Str("uri", v.URI).
				Str("route_path", v.RoutePath).
				Str("request_id", v.RequestID).
				Str("referer", v.Referer).
				Str("user_agent", v.UserAgent).
				Int("statue", v.Status).
				Int64("bytes_in", parseContentLength(v.ContentLength)).
				Int64("bytes_out", v.ResponseSize).
				Err(v.Error).
				Msg("access_request_log")

			return nil
		},
	})
}

func parseContentLength(contentLength string) int64 {
	if contentLength == "" {
		return 0
	}

	cl, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("parsing request content length")
	}

	return cl
}
