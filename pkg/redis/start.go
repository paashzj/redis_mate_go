package redis

import (
	"bufio"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"os"
	"os/exec"
	"redis_mate_go/pkg/config"
	"redis_mate_go/pkg/path"
)

func Start() {
	err := generateRedisConfig()
	if err != nil {
		logs.Error("can't start redis")
	}
	startRedis()
}

func startRedis() {
	command := exec.Command("/bin/bash", path.RedisStartScript)
	err := command.Start()
	if err != nil {
		logs.Error("start redis server failed")
	}
	command.Wait()
}

func generateRedisConfig() error {
	file, err := os.OpenFile(path.RedisConfig, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("bind " + config.RedisListenAddr))
	if err != nil {
		return err
	}
	err = writer.Flush()
	return nil
}
