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
		want      []uint
	}{
		{
			name:      "A",
			reflector: A,
			//want: decodeWiring(A.encoding()),
			want: A.ForwardWiring(),
		},
		{
			name:      "B",
			reflector: B,
			//want: decodeWiring(B.encoding()),
			want: B.ForwardWiring(),
		},
		{
			name:      "C",
			reflector: C,
			//want: decodeWiring(C.encoding()),
			want: C.ForwardWiring(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.reflector.ForwardWiring(), "ForwardWiring()")
		})
	}
}

func TestReflector_InverseWiring(t *testing.T) {
	tests := []struct {
		name   string
		wiring []uint
		want   []uint
	}{
		{
			"A",
			A.ForwardWiring(),
			A.InverseWiring(),
		},
		{
			"B",
			B.ForwardWiring(),
			B.InverseWiring()},
		{
			"C",
			C.ForwardWiring(),
			C.InverseWiring()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inverseWiring(tt.wiring)
			assert.Equal(t, tt.want, got)
		})
	}
}
