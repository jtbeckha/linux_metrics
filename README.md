# Linux Metrics

Gather metrics on your linux system and ship them off to a time-series database (influxdb) for analysis.

Off-the-shelf solutions for this exist already, but I'm having a go at learning go (:D), linux monitoring and influxdb with this project.

## Dev Environment
InfluxDB running in docker can be used to develop against

```
docker pull influxdb
docker run -p 8086:8086 \
      -v $PWD/influxdb.conf.dev:/etc/influxdb/influxdb.conf:ro \
      influxdb -config /etc/influxdb/influxdb.conf
influx -execute 'CREATE DATABASE System'
influx -execute 'SHOW DATABASES'
```
