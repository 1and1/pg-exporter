package models

import (
	"github.com/prometheus/client_golang/prometheus"
)

func newLabels(labelsKV ...string) prometheus.Labels {
	labels := prometheus.Labels{}
	var lastKey string
	for idx, arg := range labelsKV {
		if (idx % 2) == 0 {
			lastKey = arg
		} else {
			labels[lastKey] = arg
		}
	}
	return labels
}
