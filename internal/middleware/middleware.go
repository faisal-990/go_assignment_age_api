package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap" // Import Zap
)

// Change input type to *zap.Logger
func RequestLogger(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		reqID := c.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}

		c.Set("X-Request-ID", reqID)
		c.Locals("requestid", reqID)

		err := c.Next()

		duration := time.Since(start)

		// Create a slice of fields we want to log every time
		// This keeps code clean
		fields := []zap.Field{
			zap.String("request_id", reqID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", duration),
			zap.String("ip", c.IP()),
		}

		if err != nil {
			// Add the error to fields and log as Error
			fields = append(fields, zap.Error(err))
			log.Error("Request failed", fields...)
			return err
		}

		// Log as Info
		log.Info("Request processed", fields...)

		return nil
	}
}
