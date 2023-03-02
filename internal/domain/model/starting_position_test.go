package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStartingPosition(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    StartingPosition
		wantErr error
	}{
		{
			name:    "Zero",
			value:   "0",
			want:    0,
			wantErr: errors.New("'0' is an invalid enigma rotor starting position value"),
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
			wantErr: errors.New("'27' is an invalid enigma rotor starting position value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStartingPosition(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestNewStartingPositions(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    StartingPositions
		wantErr error
	}{
		{
			name:  "OK",
			value: "1,15,26",
			want: StartingPositions{
				first:  1,
				second: 15,
				third:  26,
			},
			wantErr: nil,
		},
		{
			name:    "Error 2 values",
			value:   "1,15",
			want:    StartingPositions{},
			wantErr: errors.New("please define 3 values instead of 2 at '1,15'"),
		},
		{
			name:    "Error 1 value",
			value:   "",
			want:    StartingPositions{},
			wantErr: errors.New("please define 3 values instead of 1 at ''"),
		},
		{
			name:    "Error wrong value",
			value:   "1,15,27",
			want:    StartingPositions{},
			wantErr: errors.New("'27' is an invalid enigma rotor starting position value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStartingPositions(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestStartingPositions_String(t *testing.T) {
	tests := []struct {
		name  string
		value StartingPositions
		want  string
	}{
		{
			name: "Format",
			value: StartingPositions{
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
			got := sut.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
