package timetrack

import (
	"testing"
	"time"
)

func TestTimeTrack(t *testing.T) {

	elapsed := TimeTrack(time.Now(), "test")

	if elapsed == 0 {
		t.Fail()
	}
}
