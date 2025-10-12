package main

import (
	"fmt"
	"time"
)

type TimingWheel[K comparable] struct {
	wheels      []*wheel[K]
	location    map[K]location
	tick        int64
	currentTime int64
}

type location struct {
	level     uint8
	bucketIdx uint8
}

type wheel[K comparable] struct {
	buckets       [256]map[K]int64
	currentBucket uint8
}

func NewTimingWheel[K comparable](tick time.Duration, maxTTL time.Duration) *TimingWheel[K] {

	levels := 1
	timeCapacity := tick * 256
	for timeCapacity < maxTTL && levels < 5 {
		levels++
		timeCapacity *= 256
	}

	wheels := make([]*wheel[K], levels)
	for i := range wheels {
		buckets := [256]map[K]int64{}
		for k := range buckets {
			buckets[k] = make(map[K]int64)
		}
		wheels[i] = &wheel[K]{buckets: buckets}
	}

	timingWheel := &TimingWheel[K]{
		wheels:      wheels,
		location:    make(map[K]location),
		tick:        tick.Milliseconds(),
		currentTime: time.Now().UnixMilli(),
	}

	return timingWheel
}

func (tw *TimingWheel[K]) Add(key K, expiresAt int64) {
	level, bucketIdx := tw.findLocation(expiresAt)

	tw.location[key] = location{level: level, bucketIdx: bucketIdx}

	bucket := tw.wheels[level].buckets[bucketIdx]
	bucket[key] = expiresAt
}

func (tw *TimingWheel[K]) findLocation(expiresAt int64) (uint8, uint8) {
	ticksFromNow := max((expiresAt-tw.currentTime+tw.tick)/tw.tick, 1)

	for level := 0; level < len(tw.wheels); level++ {
		levelMaxTick := int64(256) << (level * 8)
		if ticksFromNow < levelMaxTick {
			fmt.Printf("ticksFromNow %d\n", ticksFromNow)
			fmt.Printf("ticksFronNow shifted: %d\n", ticksFromNow>>(level*8))
			bucketsFromNow := uint8((ticksFromNow >> (level * 8)) & 0xff)
			fmt.Printf("bucketsFromNow %d\n", bucketsFromNow)
			bucketIdx := bucketsFromNow + tw.wheels[level].currentBucket
			fmt.Printf("bucketsIdx %d\n", bucketIdx)
			return uint8(level), bucketIdx
		}
	}
	fmt.Printf("Setting to max")
	maxLevel := uint8(len(tw.wheels) - 1)
	return maxLevel, tw.wheels[maxLevel].currentBucket - 1
}
