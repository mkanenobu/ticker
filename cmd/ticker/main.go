package main

import (
	"github.com/mkanenobu/ticker/cli"
	"github.com/mkanenobu/ticker/interval"
	"github.com/mkanenobu/ticker/ticker"
)

func main() {
	cliOptions := cli.ParseCliOptions()

	interval_, err := interval.ParseInterval(cliOptions.Interval)
	if err != nil {
		panic(err)
	}

	t := ticker.NewTicker(cliOptions.N, interval_)
	t.Start()
}
