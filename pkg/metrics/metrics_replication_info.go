package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"redis_mate_go/pkg/config"
	"redis_mate_go/pkg/util"
)

const (
	replication = "replication"
)

func metricReplicationInfo(key, value string) {
	switch key {
	case "master_repl_offset":
		replicationMasterOffset.Set(util.ConvStr2Float(value))
	}
}

var (
	replicationMasterOffset = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, replication, "master_repl_offset"),
		Help: "master_repl_offset",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)
