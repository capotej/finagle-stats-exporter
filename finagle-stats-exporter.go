package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/cactus/go-statsd-client/statsd"
	"io/ioutil"
	"log"
	"net/http"
)

type FinagleStats struct {
	Counters map[string]int
	Gauges   map[string]int
	Labels   map[string]int
	Metrics  map[string]int
}

var (
	statsd_server  = flag.String("statsd_server", "localhost:8125", "statsd server:port")
	finagle_server = flag.String("finagle_server", "localhost:9990", "finagle stats server:port")
	stats_path     = flag.String("stats_path", "stats.json", "finagle stat path")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("collecting stats from %s to %s\n", *finagle_server, *statsd_server)

	client, err := statsd.New(*statsd_server, "finagle-stats-exporter")
	if err != nil {
		log.Fatalf("Error connecting to statsd server %s", err)
	}
	defer client.Close()

	resp, err := http.Get(fmt.Sprintf("http://%s/%s", *finagle_server, *stats_path))
	if err != nil {
		log.Fatalf("Error fetching stats %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading stats %s", err)
	}

	var stats FinagleStats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		log.Fatalf("Error parsing json %s", err)
	}

	for k, v := range stats.Counters {
		err = client.Inc(k, int64(v), 1.0)
		if err != nil {
			log.Fatalf("Error sending metric: %+v\n", err)
		}
	}

	for k, v := range stats.Gauges {
		err = client.Gauge(k, int64(v), 1.0)
		if err != nil {
			fmt.Printf("Error sending metric: %+v\n", err)
		}
	}

}
