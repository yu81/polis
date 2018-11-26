package polis

import (
	"fmt"
	"testing"
)

func TestCoordinate_Distance(t *testing.T) {
	l1, l2 := NewCoordinates(135.0, 80.0), NewCoordinates(140.0, 42.0)
	fmt.Println(l1.Distance(&l2))
}

func BenchmarkLocation_Distance(b *testing.B) {
	l1, l2 := NewCoordinates(135.0, 38.0), NewCoordinates(140.0, 42.0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l1.Distance(&l2)
	}
}
