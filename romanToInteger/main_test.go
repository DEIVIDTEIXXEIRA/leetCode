package main

import (
	"testing"
)

var sync int

const test = "MCMXCIV"

// BenchmarkRomanToInt-12           2322922               518.6 ns/op             0 B/op          0 allocs/op
func BenchmarkRomanToInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sync = RomanToInt(test)
	}
}

// BenchmarkRomanToInt2-12          2452023               538.1 ns/op             0 B/op          0 allocs/op
func BenchmarkRomanToInt2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sync = RomanToInt(test)
	}
}
