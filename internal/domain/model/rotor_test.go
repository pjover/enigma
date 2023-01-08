package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotor(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Rotor
		wantErr error
	}{
		{
			name:    "I",
			value:   "i",
			want:    I,
			wantErr: nil,
		},
		{
			name:    "II",
			value:   "II",
			want:    II,
			wantErr: nil,
		},
		{
			name:    "III",
			value:   "iii",
			want:    III,
			wantErr: nil,
		},
		{
			name:    "IV",
			value:   "IV",
			want:    IV,
			wantErr: nil,
		},
		{
			name:    "V",
			value:   "v",
			want:    V,
			wantErr: nil,
		},
		{
			name:    "VI",
			value:   "VI",
			want:    VI,
			wantErr: nil,
		},
		{
			name:    "VII",
			value:   "vii",
			want:    VII,
			wantErr: nil,
		},
		{
			name:    "VIII",
			value:   "VIII",
			want:    VIII,
			wantErr: nil,
		},
		{
			name:    "Invalid rotor value",
			value:   "iX",
			want:    I,
			wantErr: errors.New("'IX' is an invalid enigma rotor value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotor(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

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
			wantErr: errors.New("'X' is an invalid enigma rotor value"),
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

func TestRotors_Format(t *testing.T) {
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
			got := sut.Format()
			assert.Equal(t, tt.want, got)
		})
	}
}
