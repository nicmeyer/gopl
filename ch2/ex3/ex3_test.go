package main

import "testing"

var u = ^uint64(0)

func BenchmarkPopCountWithoutLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(u)
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(u)
	}
}
