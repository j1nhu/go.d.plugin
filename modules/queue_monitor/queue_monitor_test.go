package queue_monitor

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueueMonitor_Charts(t *testing.T) {
	qm := New()
	require.True(t, qm.Init())

	assert.NotNil(t, qm.Charts())
}

//func TestQueueMonitor_Collect(t *testing.T) {
//	tests := map[string]struct {
//		prepare       func(t *testing.T) *QueueMonitor
//		wantCollected map[string]int64
//	}{
//		"success on valid response": {
//			prepare: prepareQueueMonitor,
//			wantCollected: map[string]int64{
//				"backend": 23,
//			},
//		},
//	}
//
//	for name, test := range tests {
//		t.Run(name, func(t *testing.T) {
//			rdb := test.prepare(t)
//			ms := rdb.Collect()
//			assert.Equal(t, test.wantCollected, ms)
//		})
//	}
//}

func prepareQueueMonitor(t *testing.T) *QueueMonitor {
	qm := New()
	return qm
}
