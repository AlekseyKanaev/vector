package vector

import (
	"reflect"

	"github.com/pkg/errors"
)

type Vector []float64

const lenError = "Both vectors must have the same length"

func ScalarMul(v Vector, mul float64) Vector {
	r := make([]float64, len(v))
	for k, _ := range v {
		r[k] = v[k] * mul
	}

	return Vector(r)
}

func ScalarDiv(v Vector, div float64) Vector {
	r := make([]float64, len(v))
	for k, _ := range v {
		r[k] = v[k] / div
	}

	return Vector(r)
}

func DotProdut(v1, v2 Vector) (float64, error) {
	if len(v1) != len(v2) {
		return 0, errors.New(lenError)
	}
	var r float64

	for k := range v1 {
		r += v1[k] * v2[k]
	}

	return r, nil
}

func CrossProduct(v1, v2 Vector) (Vector, error) {
	if len(v1) != 3 || len(v2) != 3 {
		return Vector{}, errors.New("Both vectors must be of 3 lengh")
	}

	var i, j, k float64
	i = v1[1]*v2[2] - v1[2]*v2[1]
	j = v1[2]*v2[0] - v1[0]*v2[2]
	k = v1[0]*v2[1] - v1[1]*v2[0]

	return Vector([]float64{i, j, k}), nil
}

func Mul(v1, v2 Vector) (Vector, error) {
	if len(v1) != len(v2) {
		return Vector{}, errors.New(lenError)
	}

	v := make([]float64, len(v1))
	for k := range v1 {
		v[k] = v1[k] * v2[k]
	}

	return Vector(v), nil
}

func Add(v1, v2 Vector) (Vector, error) {
	if len(v1) != len(v2) {
		return Vector{}, errors.New(lenError)
	}

	v := make([]float64, len(v1))
	for k := range v1 {
		v[k] = v1[k] + v2[k]
	}

	return Vector(v), nil
}

func Negate(v1 Vector) Vector {
	v := make([]float64, len(v1))
	for k := range v1 {
		v[k] = -v1[k]
	}

	return Vector(v)
}

func Eq(v1, v2 Vector) bool {
	if len(v1) != len(v2) {
		return false
	}

	return reflect.DeepEqual(v1, v2)
}
