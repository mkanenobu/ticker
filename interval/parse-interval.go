package interval

import "time"

func ParseInterval(interval string) (time.Duration, error) {
	return time.ParseDuration(interval)
}
