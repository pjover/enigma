package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlugboardCable(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    PlugboardCable
		wantErr error
	}{
		{
			name:    "AB",
			value:   "AB",
			want:    PlugboardCable{from: 0, to: 1},
			wantErr: nil,
		},
		{
			name:    "YZ",
			value:   "YZ",
			want:    PlugboardCable{from: 24, to: 25},
			wantErr: nil,
		},
		{
			name:    "ZY",
			value:   "ZY",
			want:    PlugboardCable{from: 25, to: 24},
			wantErr: nil,
		},
		{
			name:    "Empty input",
			value:   "",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'' is an invalid enigma plugboard cable value"),
		},
		{
			name:    "Too short",
			value:   "A",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'A' is an invalid enigma plugboard cable value"),
		},
		{
			name:    "Too long",
			value:   "ABC",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'ABC' is an invalid enigma plugboard cable value"),
		},
		{
			name:    "Lower boundary",
			value:   "@Y",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'@' is an invalid enigma plugboard value"),
		},
		{
			name:    "Upper boundary",
			value:   "A[",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'[' is an invalid enigma plugboard value"),
		},
		{
			name:    "Repeated value",
			value:   "AA",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("cannot repeat values in a enigma plugboard cable value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlugboardCable(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestNewPlugboard(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Plugboard
		wantErr error
	}{
		{
			name:  "OK",
			value: "AB,ZY,MN",
			want: Plugboard{
				{from: 0, to: 1},
				{from: 25, to: 24},
				{from: 12, to: 13},
			},
			wantErr: nil,
		},
		{
			name:    "Empty plugboard",
			value:   "",
			want:    Plugboard{},
			wantErr: nil,
		},
		{
			name:    "Wrong cable",
			value:   "AB,ZY,MM",
			want:    Plugboard{},
			wantErr: errors.New("cannot repeat values in a enigma plugboard cable value"),
		},
		{
			name:    "Repeat the same value in different cables",
			value:   "AB,ZY,AM",
			want:    Plugboard{},
			wantErr: errors.New("cannot repeat the same value 'A' in different cables"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlugboard(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestPlugboard_GetFreePlugs(t *testing.T) {
	tests := []struct {
		name string
		p    Plugboard
		want []int
	}{
		{
			"No free plugs",
			getPlugboard("AB,CD,EF,GH,IJ,KL,MN,OP,QR,ST,UV,WX,YZ"),
			[]int{},
		},
		{
			"Some free plugs",
			getPlugboard("BM,DH,RS,KN,GZ,FQ"),
			[]int{0, 2, 4, 8, 9, 11, 14, 15, 19, 20, 21, 22, 23, 24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.p.GetFreePlugs(), "GetFreePlugs()")
		})
	}
}

func getPlugboard(value string) Plugboard {
	p, _ := NewPlugboard(value)
	return p
}
