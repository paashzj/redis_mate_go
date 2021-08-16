package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"redis_mate_go/pkg/config"
	"redis_mate_go/pkg/util"
)

const (
	memory = "memory"
)

func metricMemoryInfo(key, value string) {
	switch key {
	case "used_memory_rss":
		memoryUsedRss.Set(util.ConvStr2Float(value))
	}
}

var (
	memoryUsedRss = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, memory, "used_rss"),
		Help: "used_rss",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)
