package ins

import (
	"fmt"
	"time"
)

func TimeYYmmddHHMMSS(t *time.Time) string {
	if t == nil {
		tm := time.Now()
		t = &tm;
	}

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
