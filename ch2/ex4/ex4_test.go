package main

import "testing"

var u = ^uint64(0)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(u)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(u)
	}
}
