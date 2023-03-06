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
			name:  "Empty plugboard",
			value: "",
			want: Plugboard{
				cables: []PlugboardCable{},
				wiring: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
			},
			wantErr: nil,
		},
		{
			name:  "One cable",
			value: "AB",
			want: Plugboard{
				cables: []PlugboardCable{
					{from: 0, to: 1},
				},
				wiring: []int{1, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
			},
			wantErr: nil,
		},
		{
			name:  "Two cables",
			value: "AB,CD",
			want: Plugboard{
				cables: []PlugboardCable{
					{from: 0, to: 1},
					{from: 2, to: 3},
				},
				wiring: []int{1, 0, 3, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
			},
			wantErr: nil,
		},
		{
			name:  "All cables",
			value: "AB,CD,EF,GH,IJ,KL,MN,OP,QR,ST,UV,WX,YZ",
			want: Plugboard{
				cables: []PlugboardCable{
					{from: 0, to: 1},
					{from: 2, to: 3},
					{from: 4, to: 5},
					{from: 6, to: 7},
					{from: 8, to: 9},
					{from: 10, to: 11},
					{from: 12, to: 13},
					{from: 14, to: 15},
					{from: 16, to: 17},
					{from: 18, to: 19},
					{from: 20, to: 21},
					{from: 22, to: 23},
					{from: 24, to: 25},
				},
				wiring: []int{1, 0, 3, 2, 5, 4, 7, 6, 9, 8, 11, 10, 13, 12, 15, 14, 17, 16, 19, 18, 21, 20, 23, 22, 25, 24},
			},
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

func TestPlugboard_Forward(t *testing.T) {
	p1 := getPlugboard("")
	p2 := getPlugboard("BM,DH,RS,KN,GZ,FQ")
	tests := []struct {
		name      string
		Plugboard Plugboard
		i         int
		want      int
	}{
		{
			"Without cables, 0",
			p1,
			0,
			0,
		},
		{
			"Without cables, 1",
			p1,
			1,
			1,
		},
		{
			"Without cables, 7",
			p1,
			7,
			7,
		},
		{
			"Without cables, 13",
			p1,
			13,
			13,
		},
		{
			"Without cables, 20",
			p1,
			20,
			20,
		},
		{
			"Without cables, 24",
			p1,
			24,
			24,
		},
		{
			"Without cables, 25",
			p1,
			25,
			25,
		},
		{
			"With cables, 1",
			p2,
			1,
			12,
		},
		{
			"With cables, 7",
			p2,
			7,
			3,
		},
		{
			"With cables, 13",
			p2,
			13,
			10,
		},
		{
			"With cables, 20",
			p2,
			20,
			20,
		},
		{
			"With cables, 24",
			p2,
			24,
			24,
		},
		{
			"Without cables, 25",
			p2,
			25,
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.Plugboard.Forward(tt.i)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func getPlugboard(value string) Plugboard {
	p, _ := NewPlugboard(value)
	return p
}
