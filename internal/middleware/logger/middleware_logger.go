package midLog

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func LogMiddleware(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			local_logger := logger.With("qqq", "---")
			local_logger.Debug("logger in mddleware strated")
			local_logger.Info("processing GET request")
			return next(c)
		}
	}
}
