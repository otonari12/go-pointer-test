package main

import (
	"fmt"
	"testing"
)

func BenchmarkSmallValues(b *testing.B) {
	// １回だと差がでなかったので１万回
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str := "string"
		PutStringValues(str)
	}
}

func BenchmarkSmallPointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str := "string"
		PutStringPointer(&str)
	}
}

func BenchmarkPutI5Value(b *testing.B) {
	// １回だと差がでなかったので１万回
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model := I5{}
		PutI5Value(model)
	}
}

func BenchmarkPutI5Pointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model := I5{}
		PutI5Pointer(&model)
	}
}

func BenchmarkPutI100Value(b *testing.B) {
	// １回だと差がでなかったので１万回
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model := I100{}
		PutI100Value(model)
	}
}

func BenchmarkPutI100Pointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model := I100{}
		PutI100Pointer(&model)
	}
}

func BenchmarkGetI5Value(b *testing.B) {
	var a I5
	// １回だと差がでなかったので１万回
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetI5Value()
	}
	fmt.Println(a)
}

func BenchmarkGetI5Pointer(b *testing.B) {
	var a *I5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = GetI5Pointer()
	}
	fmt.Println(a)
}
