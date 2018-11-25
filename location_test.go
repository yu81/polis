package polis

import (
	"fmt"
	"testing"
)

func TestLocation_Distance(t *testing.T) {
	l1, l2 := NewLocation(135.0, 80.0), NewLocation(140.0, 42.0)
	fmt.Println(l1.Distance(&l2))
}

func BenchmarkLocation_Distance(b *testing.B) {
	l1, l2 := NewLocation(135.0, 38.0), NewLocation(140.0, 42.0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l1.Distance(&l2)
	}
}
