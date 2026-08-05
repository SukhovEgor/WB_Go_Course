package collector

import (
	"runtime"
	"time"
)

type Metrics struct {
	Alloc uint64

	TotalAlloc uint64

	Sys uint64

	HeapAlloc uint64

	HeapSys uint64

	HeapIdle uint64

	HeapInuse uint64

	HeapReleased uint64

	HeapObjects uint64

	StackInuse uint64

	StackSys uint64

	NumGC uint32

	LastGC uint64

	PauseTotalNs uint64

	NextGC uint64

	GCCPUFraction float64

	NumGoroutine int

	CollectedAt time.Time
}

func Collect() Metrics {

	var stats runtime.MemStats

	runtime.ReadMemStats(&stats)

	return Metrics{
		Alloc: stats.Alloc,

		TotalAlloc: stats.TotalAlloc,

		Sys: stats.Sys,

		HeapAlloc: stats.HeapAlloc,

		HeapSys: stats.HeapSys,

		HeapIdle: stats.HeapIdle,

		HeapInuse: stats.HeapInuse,

		HeapReleased: stats.HeapReleased,

		HeapObjects: stats.HeapObjects,

		StackInuse: stats.StackInuse,

		StackSys: stats.StackSys,

		NumGC: stats.NumGC,

		LastGC: stats.LastGC,

		PauseTotalNs: stats.PauseTotalNs,

		NextGC: stats.NextGC,

		GCCPUFraction: stats.GCCPUFraction,

		NumGoroutine: runtime.NumGoroutine(),

		CollectedAt: time.Now(),
	}
}