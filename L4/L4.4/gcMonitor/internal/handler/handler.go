package handler

import (
	"fmt"
	"gcMonitor/internal/collector"
	"net/http"
	"runtime"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Metrics(
	w http.ResponseWriter,
	r *http.Request,
) {

	m := collector.Collect()

	w.Header().Set(
		"Content-Type",
		"text/plain; version=0.0.4",
	)

	writeGauge(w,
		"go_memory_alloc_bytes",
		"Currently allocated bytes",
		m.Alloc,
	)

	writeGauge(w,
		"go_memory_total_alloc_bytes",
		"Total allocated bytes",
		m.TotalAlloc,
	)

	writeGauge(w,
		"go_memory_sys_bytes",
		"Memory obtained from OS",
		m.Sys,
	)

	writeGauge(w,
		"go_memory_heap_alloc_bytes",
		"Heap allocated bytes",
		m.HeapAlloc,
	)

	writeGauge(w,
		"go_memory_heap_sys_bytes",
		"Heap system bytes",
		m.HeapSys,
	)

	writeGauge(w,
		"go_memory_heap_idle_bytes",
		"Heap idle bytes",
		m.HeapIdle,
	)

	writeGauge(w,
		"go_memory_heap_inuse_bytes",
		"Heap in use bytes",
		m.HeapInuse,
	)

	writeGauge(w,
		"go_memory_heap_released_bytes",
		"Released heap bytes",
		m.HeapReleased,
	)

	writeGauge(w,
		"go_memory_heap_objects",
		"Heap objects",
		m.HeapObjects,
	)

	writeGauge(w,
		"go_memory_stack_inuse_bytes",
		"Stack in use",
		m.StackInuse,
	)

	writeGauge(w,
		"go_memory_stack_sys_bytes",
		"Stack system memory",
		m.StackSys,
	)

	writeGauge(w,
		"go_gc_cycles_total",
		"Completed GC cycles",
		uint64(m.NumGC),
	)

	writeGauge(w,
		"go_gc_pause_total_ns",
		"Total GC pause",
		m.PauseTotalNs,
	)

	writeGauge(w,
		"go_gc_next_bytes",
		"Next GC target",
		m.NextGC,
	)

	writeGauge(w,
		"go_gc_last_time_unix",
		"Last GC time",
		m.LastGC,
	)

	fmt.Fprintf(w,
		"# HELP go_gc_cpu_fraction GC CPU fraction\n")

	fmt.Fprintf(w,
		"# TYPE go_gc_cpu_fraction gauge\n")

	fmt.Fprintf(
		w,
		"go_gc_cpu_fraction %f\n",
		m.GCCPUFraction,
	)

	writeGauge(w,
		"go_goroutines",
		"Current goroutines",
		uint64(m.NumGoroutine),
	)
}

func (h *Handler) RunGC(
	w http.ResponseWriter,
	r *http.Request,
) {

	runtime.GC()

	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(
		w,
		"GC completed",
	)
}

func writeGauge(
	w http.ResponseWriter,
	name string,
	help string,
	value uint64,
) {

	fmt.Fprintf(
		w,
		"# HELP %s %s\n",
		name,
		help,
	)

	fmt.Fprintf(
		w,
		"# TYPE %s gauge\n",
		name,
	)

	fmt.Fprintf(
		w,
		"%s %d\n",
		name,
		value,
	)
}