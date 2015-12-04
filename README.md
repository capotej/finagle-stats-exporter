# Finagle Stats Exporter

A simple tool that takes stats output from finagle and ships it to statsd, meant to be run from crond.

### Installation

With go:

    go get github.com/capotej/finagle-stats-exporter 

Just the binary:

    curl https://github.com/capotej/finagle-stats-exporter/releases/download/v1.0.0/finagle-stats-exporter

### Usage

    Usage of finagle-stats-exporter:
      -finagle_server="localhost:9990": finagle stats server:port
      -stats_path="stats.json": finagle stat path
      -statsd_server="localhost:8125": statsd server:port
      -metrics="true": metrics style finagle stats (non-ostrich)"
