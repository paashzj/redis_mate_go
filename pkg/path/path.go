package path

import (
	"os"
	"path/filepath"
)

// redis
var (
	RedisHome   = os.Getenv("REDIS_HOME")
	RedisConfig = filepath.FromSlash(RedisHome + "/" + "redis.conf")
)

// mate
var (
	RedisMatePath    = filepath.FromSlash(RedisHome + "/mate")
	RedisScripts     = filepath.FromSlash(RedisMatePath + "/scripts")
	RedisStartScript = filepath.FromSlash(RedisScripts + "/start-redis.sh")
)
