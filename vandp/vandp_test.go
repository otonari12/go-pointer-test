package vandp

import (
	"testing"
	"time"
)

// 値渡しのベンチマーク
func BenchmarkI1Value(b *testing.B) {
	var item = I1{0 + 4 + 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValueProcessI1(item)
	}
}

// ポインタ渡しのベンチマーク
func BenchmarkI1Pointer(b *testing.B) {
	var item = I1{0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessI1(&item)
	}
}

// ポインタ渡しのベンチマーク。ヒープ領域を利用する。
func BenchmarkI1PointerHeap(b *testing.B) {
	var item = I1{0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessHeapI1(&item)
	}
}

// 構造体をNew + 値渡しのベンチマーク
func BenchmarkI1ValueAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I1{int64(time.Now().Nanosecond())}
		ValueProcessI1(item)
	}
}

// 構造体New + ポインタ渡しのベンチマーク
func BenchmarkI1PointerAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I1{0}
		PointerProcessI1(&item)
	}
}

// 構造体New + ポインタ渡しのベンチマーク。ヒープ領域を利用する。
func BenchmarkI1PointerHeapAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I1{0}
		PointerProcessHeapI1(&item)
	}
}

func BenchmarkI5Value(b *testing.B) {
	var item = I5{0, 1, 2, 3, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValueProcessI5(item)
	}
}

func BenchmarkI5Pointer(b *testing.B) {
	var item = I5{0, 1, 2, 3, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessI5(&item)
	}
}

func BenchmarkI5PointerHeap(b *testing.B) {
	var item = I5{0, 1, 2, 3, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessHeapI5(&item)
	}
}

func BenchmarkI5ValueAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I5{0, 1, 2, 3, 4}
		ValueProcessI5(item)
	}
}

func BenchmarkI5PointerAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I5{0, 1, 2, 3, 4}
		PointerProcessI5(&item)
	}
}

func BenchmarkI5PointerHeapAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I5{0, 1, 2, 3, 4}
		PointerProcessHeapI5(&item)
	}
}

func BenchmarkI10Value(b *testing.B) {
	var item = I10{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValueProcessI10(item)
	}
}

func BenchmarkI10Pointer(b *testing.B) {
	var item = I10{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessI10(&item)
	}
}

func BenchmarkI10PointerHeap(b *testing.B) {
	var item = I10{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessHeapI10(&item)
	}
}

func BenchmarkI10ValueAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I10{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		ValueProcessI10(item)
	}
}

func BenchmarkI10PointerAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I10{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		PointerProcessI10(&item)
	}
}

func BenchmarkI10PointerHeapAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I10{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		PointerProcessHeapI10(&item)
	}
}

func BenchmarkI50Value(b *testing.B) {
	var item = I50{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValueProcessI50(item)
	}
}

func BenchmarkI50Pointer(b *testing.B) {
	var item = I50{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessI50(&item)
	}
}

func BenchmarkI50PointerHeap(b *testing.B) {
	var item = I50{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessHeapI50(&item)
	}
}

func BenchmarkI50ValueAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I50{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
		ValueProcessI50(item)
	}
}

func BenchmarkI50PointerAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I50{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
		PointerProcessI50(&item)
	}
}

func BenchmarkI50PointerHeapAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I50{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
		PointerProcessHeapI50(&item)
	}
}

func BenchmarkI100Value(b *testing.B) {
	var item = I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValueProcessI100(item)
	}
}

func BenchmarkI100Pointer(b *testing.B) {
	var item = I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessI100(&item)
	}
}

func BenchmarkI100PointerHeap(b *testing.B) {
	var item = I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessHeapI100(&item)
	}
}

func BenchmarkI100ValueAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
		ValueProcessI100(item)
	}
}

func BenchmarkI100PointerAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
		PointerProcessI100(&item)
	}
}

func BenchmarkI100PointerHeapAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = I100{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
		PointerProcessHeapI100(&item)
	}
}

func createBigStr() string {
	s := "0123456789"
	for i := 0; i < 20; i++ {
		s += s
	}
	return s
}

func BenchmarkIStrValue(b *testing.B) {
	var item = IStr{createBigStr()}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValueProcessIStr(item)
	}
}

func BenchmarkIStrPointer(b *testing.B) {
	var item = IStr{createBigStr()}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessIStr(&item)
	}
}

func BenchmarkIStrPointerHeap(b *testing.B) {
	var item = IStr{createBigStr()}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PointerProcessHeapIStr(&item)
	}
}

func BenchmarkIStrValueAndNew(b *testing.B) {
	var str = createBigStr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = IStr{str}
		ValueProcessIStr(item)
	}
}

func BenchmarkIStrPointerAndNew(b *testing.B) {
	var str = createBigStr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = IStr{str}
		PointerProcessIStr(&item)
	}
}

func BenchmarkIStrPointerHeapAndNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var item = IStr{str}
		PointerProcessHeapIStr(&item)
	}
}
