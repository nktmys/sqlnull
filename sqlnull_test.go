package sqlnull

import (
	"database/sql"
	"reflect"
	"testing"
)

type testStruct struct {
	Field1 string
	Field2 int
}

var (
	nilVal    = (any)(nil)
	zeroVal   = 0
	intVal    = 42
	strVal    = "hello"
	floatVal  = 3.14
	boolVal   = true
	structVal = testStruct{Field1: "test", Field2: 123}
	sliceVal  = []int{1, 2, 3}
	mapVal    = map[string]int{"a": 1, "b": 2}

	nilPtr       = (*any)(nil)
	zeroPtr      = &zeroVal
	intPtr       = &intVal
	strPtr       = &strVal
	floatPtr     = &floatVal
	boolPtr      = &boolVal
	structPtr    = &structVal
	slicePtr     = &sliceVal
	mapPtr       = &mapVal
	structNilPtr = (*testStruct)(nil)

	nilNull       = sql.Null[any]{Valid: false}
	zeroNull      = sql.Null[int]{Valid: true, V: zeroVal}
	intNull       = sql.Null[int]{Valid: true, V: intVal}
	strNull       = sql.Null[string]{Valid: true, V: strVal}
	floatNull     = sql.Null[float64]{Valid: true, V: floatVal}
	boolNull      = sql.Null[bool]{Valid: true, V: boolVal}
	structNull    = sql.Null[testStruct]{Valid: true, V: structVal}
	sliceNull     = sql.Null[[]int]{Valid: true, V: sliceVal}
	mapNull       = sql.Null[map[string]int]{Valid: true, V: mapVal}
	structNilNull = sql.Null[testStruct]{Valid: false}

	nilPtrNull       = sql.Null[*any]{Valid: false}
	zeroPtrNull      = sql.Null[*int]{Valid: true, V: zeroPtr}
	intPtrNull       = sql.Null[*int]{Valid: true, V: intPtr}
	strPtrNull       = sql.Null[*string]{Valid: true, V: strPtr}
	floatPtrNull     = sql.Null[*float64]{Valid: true, V: floatPtr}
	boolPtrNull      = sql.Null[*bool]{Valid: true, V: boolPtr}
	structPtrNull    = sql.Null[*testStruct]{Valid: true, V: structPtr}
	slicePtrNull     = sql.Null[*[]int]{Valid: true, V: slicePtr}
	mapPtrNull       = sql.Null[*map[string]int]{Valid: true, V: mapPtr}
	structNilPtrNull = sql.Null[*testStruct]{Valid: false}

	nilNullPtr       = (*sql.Null[any])(nil)
	structNilNullPtr = (*sql.Null[testStruct])(nil)
)

func equal(t *testing.T, name string, expected any, actual any) {
	t.Run(name, func(t *testing.T) {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %#v, actual %#v", expected, actual)
		}
	})
}

func TestFrom(t *testing.T) {
	equal(t, "nil", nilNull, From(nilVal))
	equal(t, "zero", zeroNull, From(zeroVal))
	equal(t, "int", intNull, From(intVal))
	equal(t, "string", strNull, From(strVal))
	equal(t, "float64", floatNull, From(floatVal))
	equal(t, "bool", boolNull, From(boolVal))
	equal(t, "struct", structNull, From(structVal))
	equal(t, "slice", sliceNull, From(sliceVal))
	equal(t, "map", mapNull, From(mapVal))

	equal(t, "nil pointer", nilPtrNull, From(nilPtr))
	equal(t, "pointer to zero", zeroPtrNull, From(zeroPtr))
	equal(t, "pointer to int", intPtrNull, From(intPtr))
	equal(t, "pointer to string", strPtrNull, From(strPtr))
	equal(t, "pointer to float64", floatPtrNull, From(floatPtr))
	equal(t, "pointer to bool", boolPtrNull, From(boolPtr))
	equal(t, "pointer to struct", structPtrNull, From(structPtr))
	equal(t, "pointer to slice", slicePtrNull, From(slicePtr))
	equal(t, "pointer to map", mapPtrNull, From(mapPtr))
	equal(t, "nil pointer to struct", structNilPtrNull, From(structNilPtr))
}

func TestFromPtr(t *testing.T) {
	equal(t, "nil pointer", nilNull, FromPtr(nilPtr))
	equal(t, "pointer to zero", zeroNull, FromPtr(zeroPtr))
	equal(t, "pointer to int", intNull, FromPtr(intPtr))
	equal(t, "pointer to string", strNull, FromPtr(strPtr))
	equal(t, "pointer to float64", floatNull, FromPtr(floatPtr))
	equal(t, "pointer to bool", boolNull, FromPtr(boolPtr))
	equal(t, "pointer to struct", structNull, FromPtr(structPtr))
	equal(t, "pointer to slice", sliceNull, FromPtr(slicePtr))
	equal(t, "pointer to map", mapNull, FromPtr(mapPtr))
	equal(t, "nil pointer to struct", structNilNull, FromPtr(structNilPtr))
}

func TestPtr(t *testing.T) {
	equal(t, "nil", &nilNull, Ptr(nilVal))
	equal(t, "zero", &zeroNull, Ptr(zeroVal))
	equal(t, "int", &intNull, Ptr(intVal))
	equal(t, "string", &strNull, Ptr(strVal))
	equal(t, "float64", &floatNull, Ptr(floatVal))
	equal(t, "bool", &boolNull, Ptr(boolVal))
	equal(t, "struct", &structNull, Ptr(structVal))
	equal(t, "slice", &sliceNull, Ptr(sliceVal))
	equal(t, "map", &mapNull, Ptr(mapVal))

	equal(t, "nil pointer", &nilPtrNull, Ptr(nilPtr))
	equal(t, "pointer to zero", &zeroPtrNull, Ptr(zeroPtr))
	equal(t, "pointer to int", &intPtrNull, Ptr(intPtr))
	equal(t, "pointer to string", &strPtrNull, Ptr(strPtr))
	equal(t, "pointer to float64", &floatPtrNull, Ptr(floatPtr))
	equal(t, "pointer to bool", &boolPtrNull, Ptr(boolPtr))
	equal(t, "pointer to struct", &structPtrNull, Ptr(structPtr))
	equal(t, "pointer to slice", &slicePtrNull, Ptr(slicePtr))
	equal(t, "pointer to map", &mapPtrNull, Ptr(mapPtr))
	equal(t, "nil pointer to struct", &structNilPtrNull, Ptr(structNilPtr))
}

func TestPtrOrNil(t *testing.T) {
	equal(t, "nil", nilNullPtr, PtrOrNil(nilPtr))
	equal(t, "zero", &zeroNull, PtrOrNil(zeroPtr))
	equal(t, "int", &intNull, PtrOrNil(intPtr))
	equal(t, "string", &strNull, PtrOrNil(strPtr))
	equal(t, "float64", &floatNull, PtrOrNil(floatPtr))
	equal(t, "bool", &boolNull, PtrOrNil(boolPtr))
	equal(t, "struct", &structNull, PtrOrNil(structPtr))
	equal(t, "slice", &sliceNull, PtrOrNil(slicePtr))
	equal(t, "map", &mapNull, PtrOrNil(mapPtr))
	equal(t, "nil pointer to struct", structNilNullPtr, PtrOrNil(structNilPtr))
}
