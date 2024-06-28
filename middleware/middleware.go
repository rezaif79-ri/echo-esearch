package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func SetEchoMiddleware(e *echo.Echo) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {

			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	// e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	// 	LogLatency:       true,
	// 	LogProtocol:      true,
	// 	LogRemoteIP:      true,
	// 	LogHost:          true,
	// 	LogMethod:        true,
	// 	LogURI:           true,
	// 	LogURIPath:       true,
	// 	LogRoutePath:     true,
	// 	LogRequestID:     true,
	// 	LogReferer:       true,
	// 	LogUserAgent:     true,
	// 	LogStatus:        true,
	// 	LogError:         true,
	// 	LogContentLength: true,
	// 	LogResponseSize:  true,
	// 	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
	// 		fmt.Printf("%s %s %dms, Size: %d\n", v.Method, v.URIPath, v.Latency.Milliseconds(), v.ResponseSize)
	// 		return nil
	// 	},
	// }))
}
