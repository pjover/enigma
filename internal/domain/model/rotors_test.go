package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotors(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Rotors
		wantErr error
	}{
		{
			name:  "OK",
			value: "I,VIII,II",
			want: Rotors{
				first:  I,
				second: VIII,
				third:  II,
			},
			wantErr: nil,
		},
		{
			name:    "Error 2 values",
			value:   "I,VIII",
			want:    Rotors{},
			wantErr: errors.New("please define 3 values instead of 2 at 'I,VIII'"),
		},
		{
			name:    "Error 1 value",
			value:   "",
			want:    Rotors{},
			wantErr: errors.New("please define 3 values instead of 1 at ''"),
		},
		{
			name:    "Error wrong value",
			value:   "I,II,X",
			want:    Rotors{},
			wantErr: errors.New("'X' is an invalid enigma rotor number"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotors(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRotors_String(t *testing.T) {
	tests := []struct {
		name  string
		value Rotors
		want  string
	}{
		{
			name: "Format",
			value: Rotors{
				first:  I,
				second: VIII,
				third:  II,
			},
			want: "rotors: I,VIII,II",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := tt.value
			got := sut.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
