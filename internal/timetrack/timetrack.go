package timetrack

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time, name string) time.Duration {
	elapsed := time.Since(start)
	fmt.Printf("\n-----\n%s took %s\n-----\n", name, elapsed)
	return elapsed
}
