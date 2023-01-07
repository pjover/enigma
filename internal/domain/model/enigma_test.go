package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRotors(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Rotors
		wantErr error
	}{
		{
			name:  "Happy case",
			value: "I,VIII,II",
			want: Rotors{
				first:  I,
				second: VIII,
				third:  II,
			},
			wantErr: nil,
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
