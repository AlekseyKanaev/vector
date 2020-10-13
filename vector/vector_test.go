package vector

import (
	"reflect"
	"testing"
)

func TestScalarMul(t *testing.T) {
	data := make([]float64, 3)
	data[0] = 1
	data[1] = 1
	data[2] = 1
	v := Vector(data)

	data2 := make([]float64, 3)
	data2[0] = 2
	data2[1] = 2
	data2[2] = 2
	want := Vector(data2)

	if got := ScalarMul(v, 2); !reflect.DeepEqual(got, want) {
		t.Errorf("ScalarMul() = %v, want %v", got, want)
	}
}

func TestNegate(t *testing.T) {
	type args struct {
		v1 Vector
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{
			name: "All positive",
			args: args{v1: []float64{1, 2, 3}},
			want: Vector([]float64{-1, -2, -3}),
		},
		{
			name: "All negative",
			args: args{v1: []float64{-1, -2, -3}},
			want: Vector([]float64{1, 2, 3}),
		},
		{
			name: "Mixed",
			args: args{v1: []float64{-1, 0, 3}},
			want: Vector([]float64{1, 0, -3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Negate(tt.args.v1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEq(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "All positive",
			args: args{v1: []float64{1, 2, 3}, v2: []float64{-1, -2, -3}},
			want: false,
		},
		{
			name: "All negative",
			args: args{v1: []float64{-1, -2, -3}, v2: []float64{-1, -2, -3}},
			want: true,
		},
		{
			name: "Mixed",
			args: args{v1: []float64{-1, 0, 3}, v2: []float64{-1, 0, 3}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eq(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name    string
		args    args
		want    Vector
		wantErr bool
	}{
		{
			name:    "All zeros",
			args:    args{v1: []float64{1, 2, 3}, v2: []float64{-1, -2, -3}},
			want:    []float64{0, 0, 0},
			wantErr: false,
		},
		{
			name:    "Error",
			args:    args{v1: []float64{1, 2, 3, 4}, v2: []float64{-1, -2, -3}},
			want:    []float64{0, 0, 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.v1, tt.args.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name    string
		args    args
		want    Vector
		wantErr bool
	}{
		{
			name: "Simple",
			args: args{v1: []float64{-1, 0, 3}, v2: []float64{1, 0, 2}},
			want: []float64{-1, 0, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mul(tt.args.v1, tt.args.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrossProduct(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name    string
		args    args
		want    Vector
		wantErr bool
	}{
		{
			name:    "Simple",
			args:    args{v1: []float64{3, -3, 1}, v2: []float64{4, 9, 2}},
			want:    []float64{-15, -2, 39},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CrossProduct(tt.args.v1, tt.args.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrossProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CrossProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDotProdut(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "Simple",
			args:    args{v1: []float64{3, -3, 1}, v2: []float64{4, 9, 2}},
			want:    17,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DotProdut(tt.args.v1, tt.args.v2)
			if (err != nil) != tt.wantErr {
				t.Errorf("DotProdut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DotProdut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScalarDiv(t *testing.T) {
	type args struct {
		v   Vector
		div float64
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{
			name: "Simple",
			args: args{v: []float64{2, 4, 6}, div: 2},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScalarDiv(tt.args.v, tt.args.div); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScalarDiv() = %v, want %v", got, tt.want)
			}
		})
	}
}
