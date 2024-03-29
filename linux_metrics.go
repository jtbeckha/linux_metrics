package main

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/jtbeckha/linux_metrics/cpu"
	"github.com/jtbeckha/linux_metrics/network"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	Environment   string `yaml:"environment"`
	InfluxAddress string `yaml:"influxAddress"`
}

func main() {
	configYml, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	conf := &Config{}
	if err := yaml.Unmarshal(configYml, conf); err != nil {
		panic(err)
	}

	influx, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: conf.InfluxAddress,
	})
	if err != nil {
		panic(err)
	}

	defer influx.Close()

	hostname, err := os.Hostname()
	bpConfig := client.BatchPointsConfig{
		Database:        "System",
		RetentionPolicy: "autogen",
	}
	bps, err := client.NewBatchPoints(bpConfig)
	if err != nil {
		panic(err)
	}

	networkMetrics := network.GetMetrics()
	point, err := client.NewPoint(
		"network",
		map[string]string{"hostname": hostname},
		networkMetrics,
		time.Now().UTC())
	if err != nil {
		panic(err)
	}
	bps.AddPoint(point)

	cpuMetrics := cpu.GetMetrics()
	point, err = client.NewPoint(
		"cpu",
		map[string]string{"hostname": hostname},
		cpuMetrics,
		time.Now().UTC())
	if err != nil {
		panic(err)
	}
	bps.AddPoint(point)

	err = influx.Write(bps)
	if err != nil {
		panic(err)
	}
}
