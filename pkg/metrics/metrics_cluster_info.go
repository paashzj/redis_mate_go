package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"redis_mate_go/pkg/config"
	"redis_mate_go/pkg/util"
)

const (
	cluster = "cluster"
)

func metricClusterInfo(key, value string) {
	switch key {
	case "cluster_enabled":
		clusterEnabled.Set(util.ConvStr2Float(value))
	}
}

var (
	clusterEnabled = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, cluster, "enabled"),
		Help: "cluster_enabled",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)
