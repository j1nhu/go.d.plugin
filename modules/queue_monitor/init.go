package queue_monitor

import "github.com/netdata/go.d.plugin/agent/module"

func (qm *QueueMonitor) initCharts() (*module.Charts, error) {
	return queueMonitorCharts.Copy(), nil
}
