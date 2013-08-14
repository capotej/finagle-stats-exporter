# Finagle Stats Exporter

A simple tool that takes stats output from finagle and ships it to statsd.

### Installation
    go get github.com/capotej/finagle-stats-exporter 

### Usage

    Usage of finagle-stats-exporter:
      -finagle_server="localhost:9990": finagle stats server:port
      -stats_path="stats.json": finagle stat path
      -statsd_server="localhost:8125": statsd server:port
