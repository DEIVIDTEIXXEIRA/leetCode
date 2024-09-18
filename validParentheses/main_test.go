package main

import "testing"

const test = "([{}])"

var sync bool

// BenchmarkIsValid-12     11653683               115.2 ns/op            24 B/op          2 allocs/op
func BenchmarkIsValid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sync = IsValid(test)
	}
}
