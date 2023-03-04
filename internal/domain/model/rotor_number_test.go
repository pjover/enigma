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
	type args struct {
		notchPosition int
		rotorPosition int
	}
	tests := []struct {
		name string
		r    RotorNumber
		args args
		want bool
	}{
		{
			"I, is notch",
			I,
			args{1, 1},
			true,
		},
		{
			"I, is not notch",
			I,
			args{1, 2},
			false,
		},
		{
			"II, is notch",
			II,
			args{2, 2},
			true,
		},
		{
			"II, is not notch",
			II,
			args{3, 2},
			false,
		},
		{
			"III, is notch",
			III,
			args{3, 3},
			true,
		},
		{
			"III, is not notch",
			III,
			args{3, 2},
			false,
		},
		{
			"IV, is notch",
			IV,
			args{4, 4},
			true,
		},
		{
			"IV, is not notch",
			IV,
			args{4, 5},
			false,
		},
		{
			"V, is notch",
			V,
			args{25, 25},
			true,
		},
		{
			"V, is not notch",
			V,
			args{21, 22},
			false,
		},
		{
			"VI, is notch",
			VI,
			args{6, 12},
			true,
		},
		{
			"VI, is not notch",
			VI,
			args{1, 21},
			false,
		},
		{
			"VII, is notch",
			VII,
			args{6, 25},
			true,
		},
		{
			"VII, is not notch",
			VII,
			args{1, 26},
			false,
		},
		{
			"VIII, is notch",
			VIII,
			args{6, 12},
			true,
		},
		{
			"VIII, is not notch",
			VIII,
			args{1, 13},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.r.IsAtNotch(tt.args.notchPosition, tt.args.rotorPosition), "IsAtNotch(%v, %v)", tt.args.notchPosition, tt.args.rotorPosition)
		})
	}
}

func TestRotorNumber_NotchPosition(t *testing.T) {
	tests := []struct {
		name string
		r    RotorNumber
		want uint
	}{
		{
			"I",
			I,
			16,
		},
		{
			"II",
			II,
			4,
		},
		{
			"III",
			III,
			21,
		},
		{
			"IV",
			IV,
			9,
		},
		{
			"V",
			V,
			25,
		},
		{
			"VI",
			VI,
			0,
		},
		{
			"VII",
			VII,
			0,
		},
		{
			"VII",
			VIII,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.r.NotchPosition(), "NotchPosition()")
		})
	}
}

func TestRotorNumber_ForwardWiring(t *testing.T) {
	tests := []struct {
		name     string
		encoding string
		want     []uint
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
		wiring []uint
		want   []uint
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
