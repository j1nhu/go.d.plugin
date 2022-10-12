package queue_monitor

import (
	"github.com/netdata/go.d.plugin/agent/module"
)

var backgroundQueueChartTemplate = module.Chart{
	ID:    "queue_monitor",
	Title: "Queue Length of backend queue",
	Units: "queue_length",
	Fam:   "queue_length",
	Ctx:   "queue_monitor.queue_length",
	Dims: module.Dims{
		{ID: "backend", Name: "backend"},
	},
}

var queueMonitorCharts = module.Charts{
	backgroundQueueChartTemplate.Copy(),
}
