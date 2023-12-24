package utils

func ConvertIntToInt64(i int) int64 {
	return int64(i)
}

func ConvertIntToInt64Pointer(i int) *int64 {
	x := int64(i)
	return &x
}

func ConvertIntToIntPointer(i int) *int {
	return &i
}

func ConvertIntToInt32Pointer(i int) *int32 {
	x := int32(i)
	return &x
}

func IntToBool(b int) bool {
	if b == 1 {
		return true
	}
	return false
}

func Float32PointerToFloat64Pointer(f *float32) *float64 {
	if f == nil {
		return nil
	}
	x := float64(*f)
	return &x
}

func Float32ToFloat64Pointer(f float32) *float64 {
	x := float64(f)
	return &x
}
