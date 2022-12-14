package sync_test

import (
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkSelectProdCons(b *testing.B) {
	const CallsPerSched = 1000
	procs := runtime.GOMAXPROCS(-1)
	N := int32(b.N / CallsPerSched)
	c := make(chan bool, 2*procs)
	myc := make(chan int, 128)
	myclose := make(chan bool)
	for p := 0; p < procs; p++ {
		go func() {
			// Producer: sends to myc.
			foo := 0
			// Intended to not fire during benchmarking.
			mytimer := time.After(time.Hour)
			for atomic.AddInt32(&N, -1) >= 0 {
				for g := 0; g < CallsPerSched; g++ {
					// Model some local work.
					for i := 0; i < 100; i++ {
						foo *= 2
						foo /= 2
					}
					select {
					case myc <- 1:
					case <-mytimer:
					case <-myclose:
					}
				}
			}
			myc <- 0
			c <- foo == 42
		}()
		go func() {
			// Consumer: receives from myc.
			foo := 0
			// Intended to not fire during benchmarking.
			mytimer := time.After(time.Hour)
		loop:
			for {
				select {
				case v := <-myc:
					if v == 0 {
						break loop
					}
				case <-mytimer:
				case <-myclose:
				}
				// Model some local work.
				for i := 0; i < 100; i++ {
					foo *= 2
					foo /= 2
				}
			}
			c <- foo == 42
		}()
	}
	for p := 0; p < procs; p++ {
		<-c
		<-c
	}
}

func benchmarkChanProdCons(b *testing.B, chanSize, localWork int) {
	const CallsPerSched = 1000
	procs := runtime.GOMAXPROCS(-1)
	N := int32(b.N / CallsPerSched)
	c := make(chan bool, 2*procs)
	myc := make(chan int, chanSize)
	for p := 0; p < procs; p++ {
		go func() {
			foo := 0
			for atomic.AddInt32(&N, -1) >= 0 {
				for g := 0; g < CallsPerSched; g++ {
					for i := 0; i < localWork; i++ {
						foo *= 2
						foo /= 2
					}
					myc <- 1
				}
			}
			myc <- 0
			c <- foo == 42
		}()
		go func() {
			foo := 0
			for {
				v := <-myc
				if v == 0 {
					break
				}
				for i := 0; i < localWork; i++ {
					foo *= 2
					foo /= 2
				}
			}
			c <- foo == 42
		}()
	}
	for p := 0; p < procs; p++ {
		<-c
		<-c
	}
}

func BenchmarkChanProdCons0(b *testing.B) {
	benchmarkChanProdCons(b, 0, 0)
}

func BenchmarkChanProdCons10(b *testing.B) {
	benchmarkChanProdCons(b, 10, 0)
}

func BenchmarkChanProdCons100(b *testing.B) {
	benchmarkChanProdCons(b, 100, 0)
}

func BenchmarkChanProdConsWork0(b *testing.B) {
	benchmarkChanProdCons(b, 0, 100)
}

func BenchmarkChanProdConsWork10(b *testing.B) {
	benchmarkChanProdCons(b, 10, 100)
}

func BenchmarkChanProdConsWork100(b *testing.B) {
	benchmarkChanProdCons(b, 100, 100)
}
