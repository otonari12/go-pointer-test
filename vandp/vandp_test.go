package vandp

import (
	"testing"
)

func BenchmarkV1InsertValue(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// b.Nのループの中に入れてないとheap領域が使われない
		item := I1{0}
		// heap領域の確保に時間が掛かるため、外に出す
		for n = 0; n < 10000; n++ {
			item.InsertValue()
		}
	}
}

func BenchmarkV1InsertPointer(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I1{0}
		for n = 0; n < 10000; n++ {
			item.InsertPointer()
		}
	}
}

func BenchmarkV1InsertPointerNoHeap(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I1{0}
		for n = 0; n < 10000; n++ {
			item.InsertPointerNoHeap()
		}
	}
}

func BenchmarkV5InsertValue(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I5{0, 1, 2, 3, 4}
		for n = 0; n < 10000; n++ {
			item.InsertValue()
		}
	}
}

func BenchmarkV5InsertPointer(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I5{0, 1, 2, 3, 4}
		for n = 0; n < 10000; n++ {
			item.InsertPointer()
		}
	}
}

func BenchmarkV7InsertValue(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I7{0, 1, 2, 3, 4, 5, 6}
		for n = 0; n < 10000; n++ {
			item.InsertValue()
		}
	}
}

func BenchmarkV7InsertPointer(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I7{0, 1, 2, 3, 4, 5, 6}
		for n = 0; n < 10000; n++ {
			item.InsertPointer()
		}
	}
}

func BenchmarkV8InsertValue(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I8{0, 1, 2, 3, 4, 5, 6, 7}
		for n = 0; n < 10000; n++ {
			item.InsertValue()
		}
	}
}

func BenchmarkV8InsertPointer(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I8{0, 1, 2, 3, 4, 5, 6, 7}
		for n = 0; n < 10000; n++ {
			item.InsertPointer()
		}
	}
}

func BenchmarkV9InsertValue(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I9{0, 1, 2, 3, 4, 5, 6, 7, 8}
		for n = 0; n < 10000; n++ {
			item.InsertValue()
		}
	}
}

func BenchmarkV9InsertPointer(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I9{0, 1, 2, 3, 4, 5, 6, 7, 8}
		for n = 0; n < 10000; n++ {
			item.InsertPointer()
		}
	}
}

func BenchmarkV10InsertValue(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I10{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		for n = 0; n < 10000; n++ {
			item.InsertValue()
		}
	}
}

func BenchmarkV10InsertPointer(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I10{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		for n = 0; n < 10000; n++ {
			item.InsertPointer()
		}
	}
}

func BenchmarkV100InsertValue(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
		for n = 0; n < 10000; n++ {
			item.InsertValue()
		}
	}
}

func BenchmarkV100InsertPointer(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
		for n = 0; n < 10000; n++ {
			item.InsertPointer()
		}
	}
}

func BenchmarkV100InsertPointerNoHeap(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		item := I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
		for n = 0; n < 10000; n++ {
			item.InsertPointerNoHeap()
		}
	}
}

func BenchmarkV100InsertValueNew(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n = 0; n < 10000; n++ {
			item := I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
			item.InsertValue()
		}
	}
}

func BenchmarkV100InsertPointerNew(b *testing.B) {
	n := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n = 0; n < 10000; n++ {
			item := I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
			item.InsertPointer()
		}
	}
}
