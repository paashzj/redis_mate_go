package config

import (
	"redis_mate_go/pkg/util"
)

// mate config
var (
	Cluster    string
	LogFile    string
	ConsoleLog bool
	ListenAddr string
	RemoteMode bool
)

// redis config
var (
	RedisListenAddr string
	Password        string
	Hostname        string
)

func init() {
	Cluster = util.GetEnvStr("CLUSTER", "default")
	LogFile = util.GetEnvStr("LOG_FILE", "")
	ConsoleLog = util.GetEnvBool("CONSOLE_LOG", true)
	ListenAddr = util.GetEnvStr("LISTEN_ADDR", "0.0.0.0")
	RemoteMode = util.GetEnvBool("REMOTE_MODE", true)
	RedisListenAddr = util.GetEnvStr("REDIS_LISTEN_ADDR", "0.0.0.0")
	Password = util.GetEnvStr("REMOTE_REDIS_PASSWORD", "")
	Hostname = util.GetEnvStr("REMOTE_REDIS_ADDR", "localhost:6379")
}
