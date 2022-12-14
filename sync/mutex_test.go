// package sync_test copy from sync/mutex_test.go
package sync_test

import (
	"log"
	"os"
	"sync"
	"testing"

	"github.com/felixge/fgprof"
)

func TestMain(m *testing.M) {
	if os.Getenv("fgprofile") != "" {
		f, err := os.Create(os.Getenv("fgprofile"))
		if err != nil {
			log.Fatal(err)
		}
		cancel := fgprof.Start(f, fgprof.FormatPprof)
		defer func() {
			err = cancel()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	m.Run()
}

func BenchmarkMutexUncontended(b *testing.B) {
	type PaddedMutex struct {
		sync.Mutex
		pad [128]uint8
	}
	b.RunParallel(func(pb *testing.PB) {
		var mu PaddedMutex
		for pb.Next() {
			mu.Lock()
			mu.Unlock()
		}
	})
}

func benchmarkMutex(b *testing.B, slack, work bool) {
	var mu sync.Mutex
	if slack {
		b.SetParallelism(10)
	}
	b.RunParallel(func(pb *testing.PB) {
		foo := 0
		for pb.Next() {
			mu.Lock()
			mu.Unlock()
			if work {
				for i := 0; i < 100; i++ {
					foo *= 2
					foo /= 2
				}
			}
		}
		_ = foo
	})
}

func BenchmarkMutex(b *testing.B) {
	benchmarkMutex(b, false, false)
}

func BenchmarkMutexSlack(b *testing.B) {
	benchmarkMutex(b, true, false)
}

func BenchmarkMutexWork(b *testing.B) {
	benchmarkMutex(b, false, true)
}

func BenchmarkMutexWorkSlack(b *testing.B) {
	benchmarkMutex(b, true, true)
}
