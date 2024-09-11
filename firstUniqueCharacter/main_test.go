package main

import "testing"

var sink int

const s = "abcdefghijklmnopqrstuvwxyzzyxwuvtsrqponmlkjihgfedcba"

func BenchmarkFirstChar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = FirstChar(s)
	}
}

// BenchmarkFirstChar-12             957069              1431 ns/op               0 B/op

func BenchmarkFirstUniqChar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = FirstUniqChar(s)
	}
}

// BenchmarkFirstUniqChar-12         311103              3657 ns/op            1674 B/op          5 allocs/op

func BenchmarkFirstUniqChar2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = FirstUniqChar2(s)
	}
}

// BenchmarkFirstUniqChar2-12      11360474               113.6 ns/op             0 B/op          0 allocs/op
