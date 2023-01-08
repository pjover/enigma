package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotorPosition(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    RotorPosition
		wantErr error
	}{
		{
			name:    "Zero",
			value:   "0",
			want:    0,
			wantErr: errors.New("'0' is an invalid enigma rotor position value"),
		},
		{
			name:    "One",
			value:   "1",
			want:    1,
			wantErr: nil,
		},
		{
			name:    "Twenty six",
			value:   "26",
			want:    26,
			wantErr: nil,
		},
		{
			name:    "Twenty seven",
			value:   "27",
			want:    0,
			wantErr: errors.New("'27' is an invalid enigma rotor position value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotorPosition(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestNewRotorPositions(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    RotorPositions
		wantErr error
	}{
		{
			name:  "OK",
			value: "1,15,26",
			want: RotorPositions{
				first:  1,
				second: 15,
				third:  26,
			},
			wantErr: nil,
		},
		{
			name:    "Error 2 values",
			value:   "1,15",
			want:    RotorPositions{},
			wantErr: errors.New("please define 3 values instead of 2 at '1,15'"),
		},
		{
			name:    "Error 1 value",
			value:   "",
			want:    RotorPositions{},
			wantErr: errors.New("please define 3 values instead of 1 at ''"),
		},
		{
			name:    "Error wrong value",
			value:   "1,15,27",
			want:    RotorPositions{},
			wantErr: errors.New("'27' is an invalid enigma rotor position value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotorPositions(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRotorPositions_Format(t *testing.T) {
	tests := []struct {
		name  string
		value RotorPositions
		want  string
	}{
		{
			name: "Format",
			value: RotorPositions{
				first:  1,
				second: 2,
				third:  26,
			},
			want: "pos: 1,2,26",
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
