package ndash

import (
	"sync"
	"time"
)

func Timezone8() time.Time {
	cstZone := time.FixedZone("CST", 8*3600)

	return time.Now().In(cstZone)
}

func Timezone8Format(layout ...string) string {
	//f := "2006-01-02 15:04:05"
	f := "2006/01/02 - 15:04:05"

	if len(layout) == 1 {
		f = layout[0]
	}

	return Timezone8().Format(f)
}

func SetInterval(call func(), tick <-chan time.Time) {
	go func() {
		for range tick {
			call()
		}

		var wg sync.WaitGroup
		wg.Add(1)
		wg.Wait()
	}()
}
