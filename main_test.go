package main

import (
	"testing"
)



func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sequentialSolution()
	}
}
//
//func BenchmarkSlow(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		main()
//	}
//}

