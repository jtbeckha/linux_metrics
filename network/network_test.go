package network

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

	expected := map[string]interface{}{
		"network.lo.receive.bytes":         732883748,
		"network.lo.receive.packets":       636830,
		"network.lo.receive.errs":          0,
		"network.lo.receive.drop":          0,
		"network.lo.receive.fifo":          0,
		"network.lo.receive.frame":         0,
		"network.lo.receive.compressed":    0,
		"network.lo.receive.multicast":     0,
		"network.lo.transmit.bytes":        732883748,
		"network.lo.transmit.packets":      636830,
		"network.lo.transmit.errs":         0,
		"network.lo.transmit.drop":         0,
		"network.lo.transmit.fifo":         0,
		"network.lo.transmit.colls":        0,
		"network.lo.transmit.carrier":      0,
		"network.lo.transmit.compressed":   0,
		"network.eth0.receive.bytes":       1251163119,
		"network.eth0.receive.packets":     890132,
		"network.eth0.receive.errs":        0,
		"network.eth0.receive.drop":        0,
		"network.eth0.receive.fifo":        0,
		"network.eth0.receive.frame":       0,
		"network.eth0.receive.compressed":  0,
		"network.eth0.receive.multicast":   221,
		"network.eth0.transmit.bytes":      33602049,
		"network.eth0.transmit.packets":    353730,
		"network.eth0.transmit.errs":       0,
		"network.eth0.transmit.drop":       0,
		"network.eth0.transmit.fifo":       0,
		"network.eth0.transmit.colls":      0,
		"network.eth0.transmit.carrier":    0,
		"network.eth0.transmit.compressed": 0,
	}

	actual := ParseMetrics(data)

	diffs := deep.Equal(expected, actual)
	if len(diffs) > 0 {
		t.Errorf("ParseMetrics did not return expected result, diffs=%s", diffs)
	}
}
