package benchmark

import (
	"apiProfiler/internal/service"
	"testing"
)

func BenchmarkGetUsers(b *testing.B) {

	svc := service.New()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = svc.GetUsers()
	}
}