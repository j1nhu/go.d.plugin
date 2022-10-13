package queue_monitor

import (
	"context"
	"fmt"
	"github.com/netdata/go.d.plugin/agent/module"
	"strings"
)

func (qm *QueueMonitor) collect() (map[string]int64, error) {
	collected := make(map[string]int64)

	for _, chart := range *qm.Charts() {
		qm.collectChart(collected, chart)
	}
	return collected, nil
}

func (qm *QueueMonitor) collectChart(collected map[string]int64, chart *module.Chart) {
	ctx := context.Background()
	var cursor uint64
	var n int
	for {
		var keys []string
		var err error
		keys, cursor, err = qm.rc.Scan(ctx, cursor, "queues:*", 10).Result()
		if err != nil {
			panic(err)
		}

		if len(keys) > 0 {
			n += len(keys)
			for _, key := range keys {
				colonCount := strings.Count(key, ":")
				if colonCount > 1 {
					lastInd := strings.LastIndex(key, ":")
					key = key[:lastInd]
				}

				id := fmt.Sprintf("queue_length_%s", key)
				if !qm.collectedDims[id] {
					qm.collectedDims[id] = true
					dim := &module.Dim{ID: id, Name: id}
					if err := chart.AddDim(dim); err != nil {
						qm.Warning(err)
					}
					chart.MarkNotCreated()
				}

				queueLength, err := qm.rc.LLen(ctx, key).Result()
				if err != nil {
					panic(err)
				}
				collected[id] = queueLength
			}
		}

		n += len(keys)
		if cursor == 0 {
			break
		}
	}
}
