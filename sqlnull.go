// Package sqlnull provides utilities for working with sql.Null types in Go.
package sqlnull

import (
	"database/sql"
	"reflect"
)

// From creates a sql.Null[T] value from a given value of type T.
// If the value is nil, it returns a sql.Null[T] with Valid set to false.
// If the value is a pointer and it is nil, it also returns a sql.Null[T] with Valid set to false.
// If the value is a pointer and it is valid, it returns a sql.Null[T] with V set to the pointer.
func From[T any](v T) sql.Null[T] {
	val := reflect.ValueOf(v)
	if !val.IsValid() {
		return sql.Null[T]{Valid: false}
	}
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return sql.Null[T]{Valid: false}
	}
	return sql.Null[T]{Valid: true, V: v}
}

// FromPtr creates a sql.Null[T] value from a pointer to T.
// If the pointer is nil, it returns a sql.Null[T] with Valid set to false.
// If the pointer is not nil, it dereferences the pointer and returns a sql.Null[T]
// with V set to the dereferenced value.
func FromPtr[T any](v *T) sql.Null[T] {
	if v == nil {
		return sql.Null[T]{Valid: false}
	}
	return sql.Null[T]{Valid: true, V: *v}
}

// Ptr creates a pointer to sql.Null[T] value from a given value of type T.
// If the value is nil, it returns a pointer to sql.Null[T] with Valid set to false.
// If need returns (*sql.Null[T])(nil), use PtrOrNil.
func Ptr[T any](v T) *sql.Null[T] {
	val := From(v)
	return &val
}

// PtrOrNil creates a pointer to sql.Null[T] value from a given value of type T.
// If the value is nil, it returns nil.
func PtrOrNil[T any](v *T) *sql.Null[T] {
	if v == nil {
		return nil
	}
	return Ptr(*v)
}

// ValuePtrOrNil retrieves a pointer to the value from sql.Null[T].
// If Valid is false, it returns nil.
func ValuePtrOrNil[T any](v sql.Null[T]) *T {
	if !v.Valid {
		return nil
	}
	return &v.V
}
