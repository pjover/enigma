package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRotorNumber(t *testing.T) {
	tests := []struct {
		name    string
		number  int
		want    RotorNumber
		wantErr error
	}{
		{
			name:    "Zero",
			number:  0,
			want:    0,
			wantErr: errors.New("0 is an invalid enigma rotor number"),
		},
		{
			name:    "One",
			number:  1,
			want:    I,
			wantErr: nil,
		},
		{
			name:    "Eight",
			number:  8,
			want:    VIII,
			wantErr: nil,
		},
		{
			name:    "Nine",
			number:  9,
			want:    0,
			wantErr: errors.New("9 is an invalid enigma rotor number"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotorNumber(tt.number)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRotorNumber_Encoding(t *testing.T) {
	tests := []struct {
		name string
		r    RotorNumber
		want string
	}{
		{
			"I",
			I,
			"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		},
		{
			"II",
			II,
			"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		},
		{
			"III",
			III,
			"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		},
		{
			"IV",
			IV,
			"ESOVPZJAYQUIRHXLNFTGKDCMWB",
		},
		{
			"V",
			V,
			"VZBRGITYUPSDNHLXAWMJQOFECK",
		},
		{
			"VI",
			VI,
			"JPGVOUMFYQBENHZRDKASXLICTW",
		},
		{
			"VII",
			VII,
			"NZJHGRCXMYSWBOUFAIVLPEKQDT",
		},
		{
			"VII",
			VIII,
			"FKQHTLXOCBJSPDZRAMEWNIUYGV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.r.encoding(), "encoding()")
		})
	}
}

func TestRotorNumber_IsAtNotch(t *testing.T) {
	tests := []struct {
		name            string
		r               RotorNumber
		currentPosition RotorPosition
		want            bool
	}{
		{
			"I, is notch",
			I,
			RotorPosition(16),
			true,
		},
		{
			"I, is not notch",
			I,
			RotorPosition(15),
			false,
		},
		{
			"II, is notch",
			II,
			RotorPosition(4),
			true,
		},
		{
			"II, is not notch",
			II,
			RotorPosition(5),
			false,
		},
		{
			"III, is notch",
			III,
			RotorPosition(21),
			true,
		},
		{
			"III, is not notch",
			III,
			RotorPosition(22),
			false,
		},
		{
			"IV, is notch",
			IV,
			RotorPosition(9),
			true,
		},
		{
			"IV, is not notch",
			IV,
			RotorPosition(8),
			false,
		},
		{
			"V, is notch",
			V,
			RotorPosition(25),
			true,
		},
		{
			"V, is not notch",
			V,
			RotorPosition(21),
			false,
		},
		{
			"VI, is notch",
			VI,
			RotorPosition(12),
			true,
		},
		{
			"VI, is not notch",
			VI,
			RotorPosition(24),
			false,
		},
		{
			"VII, is notch",
			VII,
			RotorPosition(25),
			true,
		},
		{
			"VII, is not notch",
			VII,
			RotorPosition(15),
			false,
		},
		{
			"VIII, is notch",
			VIII,
			RotorPosition(25),
			true,
		},
		{
			"VIII, is not notch",
			VIII,
			RotorPosition(17),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.r.IsAtNotch(tt.currentPosition)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestRotorNumber_ForwardWiring(t *testing.T) {
	tests := []struct {
		name     string
		encoding string
		want     []int
	}{
		{
			"I",
			I.encoding(),
			I.ForwardWiring(),
		},
		{
			"II",
			II.encoding(),
			II.ForwardWiring(),
		},
		{
			"III",
			III.encoding(),
			III.ForwardWiring(),
		},
		{
			"IV",
			IV.encoding(),
			IV.ForwardWiring(),
		},
		{
			"V",
			V.encoding(),
			V.ForwardWiring(),
		},
		{
			"VI",
			VI.encoding(),
			VI.ForwardWiring(),
		},
		{
			"VII",
			VII.encoding(),
			VII.ForwardWiring(),
		},
		{
			"VIII",
			VIII.encoding(),
			VIII.ForwardWiring(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, decodeWiring(tt.encoding), "decodeWiring(%v)", tt.encoding)
		})
	}
}

func TestRotorNumber_InverseWiring(t *testing.T) {
	tests := []struct {
		name   string
		wiring []int
		want   []int
	}{
		{
			"I",
			I.ForwardWiring(),
			I.InverseWiring(),
		},
		{
			"II",
			II.ForwardWiring(),
			II.InverseWiring(),
		},
		{
			"III",
			III.ForwardWiring(),
			III.InverseWiring(),
		},
		{
			"IV",
			IV.ForwardWiring(),
			IV.InverseWiring(),
		},
		{
			"V",
			V.ForwardWiring(),
			V.InverseWiring(),
		},
		{
			"VI",
			VI.ForwardWiring(),
			VI.InverseWiring(),
		},
		{
			"VII",
			VII.ForwardWiring(),
			VII.InverseWiring(),
		},
		{
			"VIII",
			VIII.ForwardWiring(),
			VIII.InverseWiring(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inverseWiring(tt.wiring)
			assert.Equal(t, tt.want, got)
		})
	}
}
