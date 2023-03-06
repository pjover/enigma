package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReflector(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		want    Reflector
		wantErr error
	}{
		{
			name:    "A",
			code:    "A",
			want:    A,
			wantErr: nil,
		},
		{
			name:    "B",
			code:    "B",
			want:    B,
			wantErr: nil,
		},
		{
			name:    "C",
			code:    "C",
			want:    C,
			wantErr: nil,
		},
		{
			name:    "error",
			code:    "a",
			want:    0,
			wantErr: errors.New("invalid reflector a"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReflector(tt.code)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestReflector_encoding(t *testing.T) {
	tests := []struct {
		name string
		r    Reflector
		want string
	}{
		{
			"A",
			A,
			"ZYXWVUTSRQPONMLKJIHGFEDCBA",
		},
		{
			"B",
			B,
			"YRUHQSLDPXNGOKMIEBFZCWVJAT",
		},
		{
			"C",
			C,
			"FVPJIAOYEDRZXWGCTKUQSBNMHL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.r.encoding(), "encoding()")
		})
	}
}

func TestReflector_ForwardWiring(t *testing.T) {
	tests := []struct {
		name      string
		reflector Reflector
		want      []int
	}{
		{
			name:      "A",
			reflector: A,
			//want: decodeWiring(A.encoding()),
			want: A.forwardWiring(),
		},
		{
			name:      "B",
			reflector: B,
			//want: decodeWiring(B.encoding()),
			want: B.forwardWiring(),
		},
		{
			name:      "C",
			reflector: C,
			//want: decodeWiring(C.encoding()),
			want: C.forwardWiring(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.reflector.forwardWiring(), "forwardWiring()")
		})
	}
}

func TestReflector_Forward(t *testing.T) {
	tests := []struct {
		name  string
		r     Reflector
		value int
		want  int
	}{
		{"A: 0", A, 0, 25},
		{"A: 1", A, 1, 24},
		{"A: 7", A, 7, 18},
		{"A: 13", A, 13, 12},
		{"A: 20", A, 20, 5},
		{"A: 24", A, 24, 1},
		{"A: 25", A, 25, 0},
		{"B: 0", B, 0, 24},
		{"B: 1", B, 1, 17},
		{"B: 7", B, 7, 3},
		{"B: 13", B, 13, 10},
		{"B: 20", B, 20, 2},
		{"B: 24", B, 24, 0},
		{"B: 25", B, 25, 19},
		{"C: 0", C, 0, 5},
		{"C: 1", C, 1, 21},
		{"C: 7", C, 7, 24},
		{"C: 13", C, 13, 22},
		{"C: 20", C, 20, 18},
		{"C: 24", C, 24, 7},
		{"C: 25", C, 25, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.r.Forward(tt.value), "Forward(%v)", tt.value)
		})
	}
}
