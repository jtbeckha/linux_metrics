package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const statsFile = "/proc/net/dev"

func GetNetworkMetrics() map[string]interface{} {
	dataBytes, err := ioutil.ReadFile(statsFile)
	if err != nil {
		log.Println("Unable to open " + statsFile + ", network stats will not be available")
		return nil
	}

	data := string(dataBytes)

	return ParseMetrics(data)
}

func ParseMetrics(data string) map[string]interface{} {
	lines := strings.Split(data, "\n")

	// "receive", "transmit"
	directions := strings.Split(lines[0], "|")[1:]
	if len(directions) != 2 {
		log.Println("Unexpected top-level header format encountered in " + statsFile)
		return nil
	}
	for index, section := range directions {
		directions[index] = strings.ToLower(strings.TrimSpace(section))
	}

	// "bytes", "packets", "errs", etc
	metricLabelsPerDirection := strings.FieldsFunc(lines[1], func(r rune) bool {
		return r == '|'
	})[1:]
	metricLabelsRx := strings.FieldsFunc(metricLabelsPerDirection[0], func(r rune) bool {
		return r == ' '
	})
	metricLabelsTx := strings.FieldsFunc(metricLabelsPerDirection[1], func(r rune) bool {
		return r == ' '
	})

	for index, section := range metricLabelsRx {
		metricLabelsRx[index] = strings.ToLower(strings.TrimSpace(section))
	}
	for index, section := range metricLabelsTx {
		metricLabelsTx[index] = strings.ToLower(strings.TrimSpace(section))
	}

	metrics := make(map[string]interface{})
	for _, line := range lines[2:] {
		if strings.TrimSpace(line) == "" {
			continue
		}

		pieces := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ':'
		})

		interfaceName := pieces[0]
		metricValues := pieces[1:]

		for index, section := range metricValues {
			metricValues[index] = strings.Trim(section, " ")
		}

		for index, value := range metricValues[:len(metricLabelsRx)] {
			direction := directions[0]
			label := metricLabelsRx[index]
			metrics[interfaceName+"."+direction+"."+label], _ = strconv.Atoi(value)

		}
		for index, value := range metricValues[len(metricLabelsTx):] {
			direction := directions[1]
			label := metricLabelsTx[index]
			metrics[interfaceName+"."+direction+"."+label], _ = strconv.Atoi(value)
		}
	}

	return metrics
}
