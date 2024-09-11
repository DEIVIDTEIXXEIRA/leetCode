package main

import "testing"

var sync bool

const test = 123456654321

// 8026539               163.9 ns/op            16 B/op          1 allocs/op
func BenchmarkIsPalindrome(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sync = IsPalindrome(test)
	}
}

// 51305547                19.62 ns/op            0 B/op          0 allocs/op
func BenchmarkIsPalindrome2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sync = IsPalindrome2(test)
	}
}
