package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRotor(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    RotorType
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
			wantErr: errors.New("'IX' is an invalid enigma rotor type value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotorType(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRotorType_Encoding(t *testing.T) {
	tests := []struct {
		name string
		r    RotorType
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
			assert.Equalf(t, tt.want, tt.r.Encoding(), "Encoding()")
		})
	}
}

func TestRotorType_IsAtNotch(t *testing.T) {
	type args struct {
		notchPosition int
		rotorPosition int
	}
	tests := []struct {
		name string
		r    RotorType
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

func TestRotorType_NotchPosition(t *testing.T) {
	tests := []struct {
		name string
		r    RotorType
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
