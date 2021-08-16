package metrics

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestRedisInfoLines(t *testing.T) {
	content, err := ioutil.ReadFile("../../doc/redis_info.txt")
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	redisInfoMetricsLines(lines)
}
