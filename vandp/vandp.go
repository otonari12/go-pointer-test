package vandp

type I1 struct {
	i0 int64
}

var num int64
var ii interface{}

func (i I1) InsertValue() {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func (i *I1) InsertPointer() {
	for j := 0; j < 1; j++ {
		ii = i
	}
}

func (i *I1) InsertPointerNoHeap() {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

type I5 struct {
	i0, i1, i2, i3, i4 int64
}

func (i I5) InsertValue() {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func (i *I5) InsertPointer() {
	for j := 0; j < 1; j++ {
		ii = i
	}
}

type I7 struct {
	i0, i1, i2, i3, i4, i5, i6 int64
}

func (i I7) InsertValue() {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func (i *I7) InsertPointer() {
	for j := 0; j < 1; j++ {
		ii = i
	}
}

type I8 struct {
	i0, i1, i2, i3, i4, i5, i6, i7 int64
}

func (i I8) InsertValue() {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func (i *I8) InsertPointer() {
	for j := 0; j < 1; j++ {
		ii = i
	}
}

type I9 struct {
	i0, i1, i2, i3, i4, i5, i6, i7, i8 int64
}

func (i I9) InsertValue() {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func (i *I9) InsertPointer() {
	for j := 0; j < 1; j++ {
		ii = i
	}
}

type I10 struct {
	i0, i1, i2, i3, i4, I1, i6, i7, i8, i9 int64
}

var I10p *I10

func (i I10) InsertValue() {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func (i *I10) InsertPointer() {
	for j := 0; j < 1; j++ {
		ii = i
	}
}

type I100 struct {
	i0, i1, i2, i3, i4, I1, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i30, i31, i32, i33, i34, i35, i36, i37, i38, i39, i40, i41, i42, i43, i44, i45, i46, i47, i48, i49, I10, I11, I12, I13, I14, I15, I16, I17, I18, I19, i60, i61, i62, i63, i64, i65, i66, i67, i68, i69, i70, i71, i72, i73, i74, i75, i76, i77, i78, i79, i80, i81, i82, i83, i84, i85, i86, i87, i88, i89, i90, i91, i92, i93, i94, i95, i96, i97, i98, i99 int64
}

func (i I100) InsertValue() {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}

func (i *I100) InsertPointer() {
	for j := 0; j < 1; j++ {
		ii = i
	}
}

func (i *I100) InsertPointerNoHeap() {
	for j := 0; j < 1; j++ {
		num = i.i0
	}
}
