package main

var StringValue string

func PutStringValues(s string) {
	StringValue = s
}

var StringPointer *string

func PutStringPointer(s *string) {
	StringPointer = s
}

type I5 struct {
	i0, i1, i2, i3, i4 int64
}

var I5Value I5

func PutI5Value(s I5) {
	I5Value = s
}

func PutI5Pointer(s *I5) {
	I5Value = *s
}

type I100 struct {
	i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i30, i31, i32, i33, i34, i35, i36, i37, i38, i39, i40, i41, i42, i43, i44, i45, i46, i47, i48, i49, i50, i51, i52, i53, i54, i55, i56, i57, i58, i59, i60, i61, i62, i63, i64, i65, i66, i67, i68, i69, i70, i71, i72, i73, i74, i75, i76, i77, i78, i79, i80, i81, i82, i83, i84, i85, i86, i87, i88, i89, i90, i91, i92, i93, i94, i95, i96, i97, i98, i99 int64
}

var I100Value I100

func PutI100Value(s I100) {
	I100Value = s
}

func PutI100Pointer(s *I100) {
	I100Value = *s
}

func GetI5Value() I5 {
	return I5{}
}

func GetI5Pointer() *I5 {
	return &I5{}
}
