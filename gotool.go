package gotool

import (
	"reflect"
	"testing"
)

//lint:ignore U1000 Ignore unused function temporarily
//goland:noinspection GoUnusedFunction
func AssertEqual[T any](t *testing.T, got T, expected T) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\n     got: %v\nexpected: %v\n", got, expected)
	}
}

//lint:ignore U1000 Ignore unused function temporarily
//goland:noinspection GoUnusedFunction
func AssertNotEqual[T any](t *testing.T, got T, expected T) {
	if reflect.DeepEqual(got, expected) {
		t.Errorf("\n     got: %v\nexpected: %v\n", got, expected)
	}
}
