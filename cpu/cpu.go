// Package cpu contains functions for gathering cpu-related metrics.
package cpu

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const metricsFile = "/proc/stat"
const namespace = "cpu"

var metricLabels = [10]string{
	"user",
	"nice",
	"system",
	"idle",
	"iowait",
	"irq",
	"softirq",
	"steal",
	"guest",
	"guest_nice",
}

// Get cpu metrics.
func GetMetrics() map[string]interface{} {
	dataBytes, err := ioutil.ReadFile(metricsFile)
	if err != nil {
		log.Println("Unable to open "+metricsFile+", network stats will not be available", err)
		return nil
	}

	data := string(dataBytes)

	return ParseMetrics(data)
}

/*
Parse cpu-related metrics from the provided data string. Data is assumed to be in the format provided by /proc/stat.
Return a map of metric name->value pairs.
*/
func ParseMetrics(data string) map[string]interface{} {
	lines := strings.Split(data, "\n")

	var aggregateCpuLine = ""
	for _, line := range lines {
		matched, err := regexp.MatchString("cpu\\s", line)
		if err != nil {
			log.Println("Unexpected error occurred matching line, ignoring", err)
		}

		if matched {
			aggregateCpuLine = line
			break
		}
	}

	if aggregateCpuLine == "" {
		log.Println("Unable to parse aggregate cpu metrics in " + metricsFile)
		return nil
	}

	metricValues := strings.FieldsFunc(aggregateCpuLine, func(r rune) bool {
		return r == ' '
	})[1:]

	if len(metricValues) > len(metricLabels) {
		log.Println(
			"Parsed more metrics than expected, metrics after " + metricLabels[len(metricLabels)-1] + " will be ignored")
	}

	metricCount := len(metricLabels)
	if metricCount > len(metricValues) {
		metricCount = len(metricValues)
	}

	metrics := make(map[string]interface{})
	for index := 0; index < metricCount; index++ {
		metricLabel := metricLabels[index]
		metricValue, _ := strconv.Atoi(metricValues[index])

		metrics[namespace+".aggregate."+metricLabel] = metricValue
	}

	return metrics
}
