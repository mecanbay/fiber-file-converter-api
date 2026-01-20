package metrics

import (
	"fiber-file-converter-api/internal/observability/metrics"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (h *Handler) Metrics() http.Handler {
	handler := promhttp.HandlerFor(
		metrics.Registry,
		promhttp.HandlerOpts{},
	)
	return handler
}
