package cli

import (
	"flag"
)

type CliOptions struct {
	N        int64
	Interval string
}

func ParseCliOptions() *CliOptions {
	var (
		n        int64
		interval string
	)

	flag.Int64Var(&n, "n", 10, "Number of ticks")
	flag.StringVar(&interval, "interval", "1s", "Interval between ticks")
	flag.Parse()

	op := &CliOptions{
		N:        n,
		Interval: interval,
	}

	return op
}
