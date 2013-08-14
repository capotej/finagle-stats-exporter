package main

import (
	"encoding/json"
	"flag"
	"fmt"
	//"github.com/cactus/go-statsd-client/statsd"
	"io/ioutil"
	"net/http"
)

type FinagleStats struct {
	Counters map[string]int
	Gauges   map[string]int
	Labels   map[string]int
	Metrics  map[string]int
}

func main() {
	var statsd_server = flag.String("statsd_server", "localhost:8125", "statsd server \"localhost:8125\"")
	var finagle_server = flag.String("finagle_server", "localhost:9990", "finagle server \"localhost:8888\"")

	fmt.Printf("collecting stats from %s to %s\n", *finagle_server, *statsd_server)

	resp, err := http.Get(fmt.Sprintf("http://%s/stats.json", *finagle_server))
	if err != nil {
		fmt.Printf("error fetching stats %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading stats %s", err)
	}

	var stats FinagleStats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		fmt.Printf("error reading json %s", err)
	}

}
