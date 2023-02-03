package metrics

import (
	"testing"

	"github.com/m-lab/go/prometheusx/promtest"
)

func TestLintMetrics(t *testing.T) {
	RequestsTotal.WithLabelValues("type", "condition", "status")
	AppEngineTotal.WithLabelValues("country")
	CurrentHeartbeatConnections.WithLabelValues("experiment").Set(0)
	LocateHealthStatus.WithLabelValues("experiment").Set(0)
	ImportMemorystoreTotal.WithLabelValues("status")
	PrometheusHealthCollectionDuration.WithLabelValues("code")
	PortChecksTotal.WithLabelValues("status")
	KubernetesRequestsTotal.WithLabelValues("status")
	KubernetesRequestTimeHistogram.WithLabelValues("healthy")
	promtest.LintMetrics(nil)
}
