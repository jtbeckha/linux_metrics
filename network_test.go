package main

import (
	"github.com/go-test/deep"
	"testing"
)

func TestParseMetrics(t *testing.T) {
	data :=
		"Inter-|   Receive                                                |  Transmit\n" +
		"face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed\n" +
		"lo: 732883748  636830    0    0    0     0          0         0 732883748  636830    0    0    0     0       0          0\n" +
		"eth0: 1251163119  890132    0    0    0     0          0       221 33602049  353730    0    0    0     0       0          0\n"

	expected := map[string]interface{} {
		"lo.receive.bytes": 732883748,
		"lo.receive.packets": 636830,
		"lo.receive.errs": 0,
		"lo.receive.drop": 0,
		"lo.receive.fifo": 0,
		"lo.receive.frame": 0,
		"lo.receive.compressed": 0,
		"lo.receive.multicast": 0,
		"lo.transmit.bytes": 732883748,
		"lo.transmit.packets": 636830,
		"lo.transmit.errs": 0,
		"lo.transmit.drop": 0,
		"lo.transmit.fifo": 0,
		"lo.transmit.colls": 0,
		"lo.transmit.carrier": 0,
		"lo.transmit.compressed": 0,
		"eth0.receive.bytes": 1251163119,
		"eth0.receive.packets": 890132,
		"eth0.receive.errs": 0,
		"eth0.receive.drop": 0,
		"eth0.receive.fifo": 0,
		"eth0.receive.frame": 0,
		"eth0.receive.compressed": 0,
		"eth0.receive.multicast": 221,
		"eth0.transmit.bytes": 33602049,
		"eth0.transmit.packets": 353730,
		"eth0.transmit.errs": 0,
		"eth0.transmit.drop": 0,
		"eth0.transmit.fifo": 0,
		"eth0.transmit.colls": 0,
		"eth0.transmit.carrier": 0,
		"eth0.transmit.compressed": 0,
	}


	actual := parseMetrics(data)

	diffs := deep.Equal(expected, actual)
	if len(diffs) > 0 {
		t.Errorf("parseMetrics did not return expected result, diffs=%s", diffs)
	}
}
