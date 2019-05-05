package main

import (
	"context"
	"github.com/influxdata/influxdb-client-go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Config struct {
	Environment 	string 	`yaml:"environment"`
	InfluxAddress 	string 	`yaml:"influxAddress"`
	InfluxToken		string	`yaml:"influxToken"`
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

	influx, err := influxdb.New(
		nil, influxdb.WithAddress(conf.InfluxAddress), influxdb.WithToken(conf.InfluxToken))
	if err != nil {
		panic(err)
	}

	hostname, err := os.Hostname()

	// TODO get actual metrics
	myMetrics := []influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{"memory": 1000, "cpu": 0.93},
			"system-metrics",
			map[string]string{"hostname": hostname},
			time.Now().UTC()),
	}

	if err := influx.Write(context.Background(), "my-awesome-bucket", "my-awesome-org", myMetrics...); err != nil {
		log.Fatal(err)
	}
}

func reportMetrics(influx *influxdb.Client, bucket string, org string, metrics []influxdb.Metric) error {
	return influx.Write(context.Background(), bucket, org, metrics...)
}
