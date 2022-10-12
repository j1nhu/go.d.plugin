package queue_monitor

import (
	"github.com/go-redis/redis/v8"
	"github.com/netdata/go.d.plugin/agent/module"
)

func (qm *QueueMonitor) initCharts() (*module.Charts, error) {
	return queueMonitorCharts.Copy(), nil
}

func (qm *QueueMonitor) initRedisClient() (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}), nil
}
