package queue_monitor

func (qm *QueueMonitor) collect() (map[string]int64, error) {
	//queueLength, err := qm.rc.LLen(context.Background(), "queues:backend").Result()
	//if err != nil {
	//	panic(err)
	//}

	collected := make(map[string]int64)
	collected["backend"] = 23
	return collected, nil
}
