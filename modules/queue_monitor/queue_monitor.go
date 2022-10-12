package queue_monitor

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/netdata/go.d.plugin/agent/module"
)

func init() {
	module.Register("queue_monitor", module.Creator{
		Create: func() module.Module { return &QueueMonitor{} }},
	)
}

type QueueMonitor struct {
	module.Base
	rc     redisClient
	charts *module.Charts
}

type redisClient interface {
	LLen(ctx context.Context, key string) *redis.IntCmd
	Ping(context.Context) *redis.StatusCmd
	Close() error
}

func (qm *QueueMonitor) Init() bool {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	pingResult, err := rdb.Ping(ctx).Result()
	fmt.Println("ping result is", pingResult)
	if err != nil {
		return false
	}
	qm.rc = rdb
	charts, err := qm.initCharts()
	if err != nil {
		qm.Errorf("init charts: %v", err)
		return false
	}
	qm.charts = charts
	return true
}

func (qm *QueueMonitor) Check() bool { return true }

func (qm *QueueMonitor) Charts() *module.Charts {
	return qm.charts
}

func (qm *QueueMonitor) Collect() map[string]int64 {
	result, err := qm.collect()
	if err != nil {
		qm.Error(err)
	}
	return result
}

func (qm *QueueMonitor) Cleanup() {
	if qm.rc == nil {
		return
	}
	err := qm.rc.Close()
	if err != nil {
		qm.Warningf("cleanup: error on closing redis client: %v", err)
	}
	qm.rc = nil
}
