package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"redis_mate_go/pkg/config"
	"redis_mate_go/pkg/util"
)

const (
	stats = "stats"
)

func metricStatsInfo(key, value string) {
	switch key {
	case "rejected_connections":
		statsRejectedConnections.Set(util.ConvStr2Float(value))
	case "instantaneous_ops_per_sec":
		statsOpsPerSec.Set(util.ConvStr2Float(value))
	case "keyspace_hits":
		statsKeyspaceHit.Set(util.ConvStr2Float(value))
	case "keyspace_misses":
		statsKeyspaceMisses.Set(util.ConvStr2Float(value))
	}
}

var (
	statsRejectedConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, stats, "rejected_connections"),
		Help: "rejected_connections",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	statsOpsPerSec = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, stats, "ops_per_sec"),
		Help: "ops_per_sec",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	statsKeyspaceHit = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, stats, "keyspace_hits"),
		Help: "keyspace_hits",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	statsKeyspaceMisses = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, stats, "keyspace_misses"),
		Help: "keyspace_misses",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)
