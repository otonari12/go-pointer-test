package vandp

import (
	"testing"
)

var item *I5
func BenchmarkGetI5Value(b *testing.B) {
	item = &I5{0}
	// １回だと差がでなかったので１万回
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item.FromValue()
	}
}

func BenchmarkGetI5Pointer2(b *testing.B) {
	item = &I5{0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item.FromPointer()
	}
}
