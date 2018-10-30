package main

import (
	"sort"
)

// ZoneLabelStats type
type ZoneLabelStats struct {
	pos     int
	rotated bool
	log     []string
	in      chan string
	out     chan []string
	reset   chan bool
	close   chan bool
}

// LabelStats type
type LabelStats []labelStat

// Len func
func (s LabelStats) Len() int { return len(s) }

// Swap func
func (s LabelStats) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type labelStatsByCount struct{ LabelStats }

func (s labelStatsByCount) Less(i, j int) bool { return s.LabelStats[i].Count > s.LabelStats[j].Count }

type labelStat struct {
	Label string
	Count int
}

// NewZoneLabelStats func
func NewZoneLabelStats(size int) *ZoneLabelStats {
	zs := &ZoneLabelStats{
		log:   make([]string, size),
		in:    make(chan string, 100),
		out:   make(chan []string),
		reset: make(chan bool),
		close: make(chan bool),
	}
	go zs.receiver()
	return zs
}

func (zs *ZoneLabelStats) receiver() {

	for {
		select {
		case new := <-zs.in:
			zs.add(new)
		case zs.out <- zs.log:
		case <-zs.reset:
			zs.pos = 0
			zs.log = make([]string, len(zs.log))
			zs.rotated = false
		case <-zs.close:
			close(zs.in)
			return
		}
	}

}

// Close func
func (zs *ZoneLabelStats) Close() {
	zs.close <- true
}

// Reset func
func (zs *ZoneLabelStats) Reset() {
	zs.reset <- true
}

// Add func
func (zs *ZoneLabelStats) Add(l string) {
	zs.in <- l
}

func (zs *ZoneLabelStats) add(l string) {
	zs.log[zs.pos] = l
	zs.pos++
	if zs.pos+1 > len(zs.log) {
		zs.rotated = true
		zs.pos = 0
	}
}

// TopCounts func
func (zs *ZoneLabelStats) TopCounts(n int) LabelStats {
	cm := zs.Counts()
	top := make(LabelStats, len(cm))
	i := 0
	for l, c := range cm {
		top[i] = labelStat{l, c}
		i++
	}
	sort.Sort(labelStatsByCount{top})
	if len(top) > n {
		others := 0
		for _, t := range top[n:] {
			others += t.Count
		}
		top = append(top[:n], labelStat{"Others", others})
	}
	return top
}

// Counts func
func (zs *ZoneLabelStats) Counts() map[string]int {
	log := (<-zs.out)

	counts := make(map[string]int)
	for i, l := range log {
		if zs.rotated == false && i >= zs.pos {
			break
		}
		counts[l]++
	}
	return counts
}
