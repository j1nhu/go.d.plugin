package queue_monitor

import (
	"context"
)

func (qm *QueueMonitor) collect() (map[string]int64, error) {
	queueLength, err := qm.rc.LLen(context.Background(), "queues:backend").Result()
	if err != nil {
		panic(err)
	}
	//ctx := context.Background()
	//
	//rdb := redis.NewClient(&redis.Options{
	//	Addr: "127.0.0.1:6379",
	//})
	//queueLength, err := rdb.LLen(ctx, "queues:backend").Result()
	//if err != nil {
	//	panic(err)
	//}

	return map[string]int64{
		"backend": queueLength,
	}, nil
}
