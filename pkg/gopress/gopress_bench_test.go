package gopress

import "testing"

func BenchmarkPush(b *testing.B) {
	gp := New(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gp.Push("job-payload")
	}
}

func BenchmarkPushPop(b *testing.B) {
	gp := New(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gp.Push("job-payload")
		gp.Pop()
	}
}
