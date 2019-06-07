package vandp

type I1 struct {
	i0 int64
}

type I5 struct {
	i0, i1, i2, i3, i4 int64
}

type I10 struct {
	i0, i1, i2, i3, i4, I1, i6, i7, i8, i9 int64
}

type I50 struct {
	i0, i1, i2, i3, i4, I1, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i30, i31, i32, i33, i34, i35, i36, i37, i38, i39, i40, i41, i42, i43, i44, i45, i46, i47, i48, i49 int64
}

type I100 struct {
	i0, i1, i2, i3, i4, I1, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i30, i31, i32, i33, i34, i35, i36, i37, i38, i39, i40, i41, i42, i43, i44, i45, i46, i47, i48, i49, I10, I11, I12, I13, I14, I15, I16, I17, I18, I19, i60, i61, i62, i63, i64, i65, i66, i67, i68, i69, i70, i71, i72, i73, i74, i75, i76, i77, i78, i79, i80, i81, i82, i83, i84, i85, i86, i87, i88, i89, i90, i91, i92, i93, i94, i95, i96, i97, i98, i99 int64
}

type IStr struct {
	i0 string
}

var num int64
var str string
var intf interface{}

func ValueProcessI1(i I1) {
	// for文を含めることで、インラインで実行されるのを防ぐ
	for j := 0; j < 1; j++ {
		// PointerProcessHeapI1参照と同等負荷の処理
		num = i.i0
	}
}

func PointerProcessI1(i *I1) {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func PointerProcessHeapI1(i *I1) {
	for j := 0; j < 1; j++ {
		// パッケージスコープの変数に代入することで、ヒープ領域を使わせる
		intf = i
	}
}

func ValueProcessI5(i I5) {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func PointerProcessI5(i *I5) {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func PointerProcessHeapI5(i *I5) {
	for j := 0; j < 1; j++ {
		intf = i
	}
}

func ValueProcessI10(i I10) {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func PointerProcessI10(i *I10) {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func PointerProcessHeapI10(i *I10) {
	for j := 0; j < 1; j++ {
		intf = i
	}
}

func ValueProcessI50(i I50) {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func PointerProcessI50(i *I50) {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func PointerProcessHeapI50(i *I50) {
	for j := 0; j < 1; j++ {
		intf = i
	}
}

func ValueProcessI100(i I100) {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func PointerProcessI100(i *I100) {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func PointerProcessHeapI100(i *I100) {
	for j := 0; j < 1; j++ {
		intf = i
	}
}

func ValueProcessIStr(i IStr) {
	for j := 0; j < 1; j++ {
		str = i.i0
	}
}

func PointerProcessIStr(i *IStr) {
	for j := 0; j < 1; j++ {
		str = i.i0
	}
}

func PointerProcessHeapIStr(i *IStr) {
	for j := 0; j < 1; j++ {
		intf = i
	}
}
