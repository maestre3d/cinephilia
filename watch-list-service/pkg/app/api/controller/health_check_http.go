package controller

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// HealthCheckHTTP handles live pulsing for any service monitoring task
//	Contains K8s default health checks
//	@API
//	@Controller
//	@Handler
//	@HTTP
type HealthCheckHTTP struct {
	readinessMu sync.RWMutex
	healthy     bool
}

func NewHealthCheckHTTP(app *fiber.App) *HealthCheckHTTP {
	health := new(HealthCheckHTTP)
	health.initRouting(app)

	// healthCheck will report the server is unhealthy for 10 seconds after
	// startup, and as healthy henceforth. Check the /healthz/readiness
	// HTTP path to see readiness.
	time.AfterFunc(10*time.Second, func() {
		health.readinessMu.Lock()
		defer health.readinessMu.Unlock()
		health.healthy = true
	})
	return health
}

func (h *HealthCheckHTTP) initRouting(router fiber.Router) {
	router.Get("/health", h.getOverall)
	router.Get("/healthz/liveness", h.getLiveness)
	router.Get("/healthz/readiness", h.getReadiness)
}

func (h *HealthCheckHTTP) getOverall(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.JSON(fiber.Map{
		"service": "watch_list",
		"status":  "ok",
	})
}

func (h *HealthCheckHTTP) getLiveness(c *fiber.Ctx) (err error) {
	defer func(started time.Time) {
		if duration := time.Now().Sub(started); duration > 10 {
			err = c.Send([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
			return
		}

		c.Status(http.StatusOK)
		err = c.Send([]byte("ok"))
	}(time.Now())

	return err
}

func (h *HealthCheckHTTP) getReadiness(c *fiber.Ctx) error {
	h.readinessMu.RLock()
	defer h.readinessMu.RUnlock()
	if !h.healthy {
		return errors.New("watch_list service not ready yet")
	}

	return c.Send([]byte("ok"))
}
