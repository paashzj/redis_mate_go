package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"redis_mate_go/pkg/config"
	"redis_mate_go/pkg/util"
)

const (
	clients = "clients"
)

func metricClientInfo(key, value string) {
	switch key {
	case "connected_clients":
		clientAbortedClients.Set(util.ConvStr2Float(value))
	}
}

var (
	clientAbortedClients = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, clients, "connected_clients"),
		Help: "connected_clients",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)
