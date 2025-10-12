package main

import (
	"testing"
	"testing/synctest"
	"time"
)

func Test_Wheel(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		now := time.Now()
		tick := 10 * time.Millisecond
		maxTTL := 300 * tick
		tw := NewTimingWheel[int](tick, maxTTL)

		if got, want := len(tw.wheels), 2; got != want {
			t.Errorf("got %d levels, want %d", got, want)
		}

		tw.Add(0, now.Add(1*tick).UnixMilli())
		if got, want := tw.location[0], (location{level: 0, bucketIdx: 2}); got != want {
			t.Errorf("got %+v location, want %+v", got, want)
		}

		tw.Add(1, now.Add(100*tick).UnixMilli())
		if got, want := tw.location[1], (location{level: 0, bucketIdx: 101}); got != want {
			t.Errorf("got %+v location, want %+v", got, want)
		}

		tw.Add(2, now.Add(255*tick).UnixMilli())
		if got, want := tw.location[2], (location{level: 1, bucketIdx: 1}); got != want {
			t.Errorf("got %+v location, want %+v", got, want)
		}

		tw.Add(3, now.Add(2*256*tick).UnixMilli())
		if got, want := tw.location[3], (location{level: 1, bucketIdx: 2}); got != want {
			t.Errorf("got %+v location, want %+v", got, want)
		}

		tw.Add(4, now.Add(10*256*256*tick).UnixMilli())
		if got, want := tw.location[4], (location{level: 1, bucketIdx: 255}); got != want {
			t.Errorf("got %+v location, want %+v", got, want)
		}

	})
}
