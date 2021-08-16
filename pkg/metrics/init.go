package metrics

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"redis_mate_go/pkg/config"
	"redis_mate_go/pkg/redis"
	"redis_mate_go/pkg/util"
	"strings"
	"time"
)

const (
	namespace = "redis"
)

func Init() {
	util.Schedule(30, time.Second, redisInfoMetrics)
}

func redisInfoMetrics() {
	lines := strings.Split(redis.Info(), "\n")
	redisInfoMetricsLines(lines)
}

func redisInfoMetricsLines(lines []string) {
	fieldClass := ""

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			logs.Debug("line is empty")
			continue
		}
		logs.Debug("line is %s", line)
		if len(line) > 0 && strings.HasPrefix(line, "# ") {
			fieldClass = line[2:]
			continue
		}

		if (len(line) < 2) || (!strings.Contains(line, ":")) {
			continue
		}
		split := strings.SplitN(line, ":", 2)
		fieldKey := split[0]
		fieldValue := split[1]

		switch fieldClass {
		case "Clients":
			metricClientInfo(fieldKey, fieldValue)
		case "Cluster":
			metricClusterInfo(fieldKey, fieldValue)
		case "Replication":
			metricReplicationInfo(fieldKey, fieldValue)
		case "Stats":
			metricStatsInfo(fieldKey, fieldValue)
		case "Memory":
			metricMemoryInfo(fieldKey, fieldValue)
		}
	}
	upGauge.Set(1)
}

var (
	upGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "redis_up",
		Help: "is redis running",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)
