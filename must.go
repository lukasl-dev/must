package must

func Boolean(b bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return b
}

func Uint8(i uint8, err error) uint8 {
	if err != nil {
		panic(err)
	}
	return i
}

func Uint16(i uint16, err error) uint16 {
	if err != nil {
		panic(err)
	}
	return i
}

func Uint32(i uint32, err error) uint32 {
	if err != nil {
		panic(err)
	}
	return i
}

func Int8(i int8, err error) int8 {
	if err != nil {
		panic(err)
	}
	return i
}

func Int16(i int16, err error) int16 {
	if err != nil {
		panic(err)
	}
	return i
}

func Int32(i int32, err error) int32 {
	if err != nil {
		panic(err)
	}
	return i
}

func Int64(i int64, err error) int64 {
	if err != nil {
		panic(err)
	}
	return i
}

func Float32(f float32, err error) float32 {
	if err != nil {
		panic(err)
	}
	return f
}

func Float64(f float64, err error) float64 {
	if err != nil {
		panic(err)
	}
	return f
}

func Complex64(c complex64, err error) complex64 {
	if err != nil {
		panic(err)
	}
	return c
}

func Complex128(c complex128, err error) complex128 {
	if err != nil {
		panic(err)
	}
	return c
}

func String(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}

func Int(i int, err error) int {
	if err != nil {
		panic(err)
	}
	return i
}

func Uint(i uint, err error) uint {
	if err != nil {
		panic(err)
	}
	return i
}

func Uintptr(i uintptr, err error) uintptr {
	if err != nil {
		panic(err)
	}
	return i
}

func Byte(b byte, err error) byte {
	if err != nil {
		panic(err)
	}
	return b
}

func Rune(r rune, err error) rune {
	if err != nil {
		panic(err)
	}
	return r
}

func Interface(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
