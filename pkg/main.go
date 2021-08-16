package main

import (
	"flag"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"redis_mate_go/pkg/config"
	"redis_mate_go/pkg/metrics"
	"redis_mate_go/pkg/redis"
)

func main() {
	logs.EnableFuncCallDepth(true)
	flag.Parse()
	if config.ConsoleLog {
		logs.SetLogger(logs.AdapterConsole)
	}
	if config.LogFile != "" {
		strings := `{"filename":"` + config.LogFile + `", "level":6}`
		logs.SetLogger(logs.AdapterFile, strings)
	}
	if !config.RemoteMode {
		logs.Info("not remote mode, start redis server")
		redis.Start()
	} else {
		logs.Info("remote mode")
	}
	logs.Debug("this is debug message")
	logs.Info("this is info message")
	logs.Error("this is error message")
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	http.Handle("/", r)

	metrics.Init()
	err := http.ListenAndServe(config.ListenAddr+":31003", nil)
	if err != nil {
		logs.Error("open http server failed")
		return
	}
	logs.Info("redis mate started")
}
