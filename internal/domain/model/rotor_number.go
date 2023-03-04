package model

import (
	"fmt"
)

type RotorNumber uint

const (
	I RotorNumber = iota
	II
	III
	IV
	V
	VI
	VII
	VIII
)

func NewRotorNumber(number int) (RotorNumber, error) {

	if number < 1 || number > 8 {
		return RotorNumber(0), fmt.Errorf("%d is an invalid enigma rotor number", number)
	}
	return RotorNumber(number - 1), nil
}

var stringValues = []string{
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"VIII",
}

func (r RotorNumber) String() string {
	return stringValues[r]
}

func firstNotchTypeFunc(notchPosition int, rotorPosition int) bool {
	return notchPosition == rotorPosition
}

func secondNotchTypeFunc(_ int, rotorPosition int) bool {
	return rotorPosition == 12 || rotorPosition == 25
}

var isAtNotchFunctions = []func(int, int) bool{
	firstNotchTypeFunc,
	firstNotchTypeFunc,
	firstNotchTypeFunc,
	firstNotchTypeFunc,
	firstNotchTypeFunc,
	secondNotchTypeFunc,
	secondNotchTypeFunc,
	secondNotchTypeFunc,
}

func (r RotorNumber) IsAtNotch(notchPosition int, rotorPosition int) bool {
	return isAtNotchFunctions[r](notchPosition, rotorPosition)
}

var notchPositionValues = []uint{
	16,
	4,
	21,
	9,
	25,
	0,
	0,
	0,
}

func (r RotorNumber) NotchPosition() uint {
	return notchPositionValues[r]
}

var rotorEncodingValues = []string{
	"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	"AJDKSIRUXBLHWTMCQGZNPYFVOE",
	"BDFHJLCPRTXVZNYEIWGAKMUSQO",
	"ESOVPZJAYQUIRHXLNFTGKDCMWB",
	"VZBRGITYUPSDNHLXAWMJQOFECK",
	"JPGVOUMFYQBENHZRDKASXLICTW",
	"NZJHGRCXMYSWBOUFAIVLPEKQDT",
	"FKQHTLXOCBJSPDZRAMEWNIUYGV",
}

func (r RotorNumber) encoding() string {
	return rotorEncodingValues[r]
}

func decodeWiring(encoding string) []uint {
	charWiring := []rune(encoding)
	wiring := make([]uint, len(charWiring))
	for i, c := range charWiring {
		wiring[i] = uint(c - 'A')
	}
	return wiring
}

var rotorForwardWiringValues = [][]uint{
	{0x4, 0xa, 0xc, 0x5, 0xb, 0x6, 0x3, 0x10, 0x15, 0x19, 0xd, 0x13, 0xe, 0x16, 0x18, 0x7, 0x17, 0x14, 0x12, 0xf, 0x0, 0x8, 0x1, 0x11, 0x2, 0x9},
	{0x0, 0x9, 0x3, 0xa, 0x12, 0x8, 0x11, 0x14, 0x17, 0x1, 0xb, 0x7, 0x16, 0x13, 0xc, 0x2, 0x10, 0x6, 0x19, 0xd, 0xf, 0x18, 0x5, 0x15, 0xe, 0x4},
	{0x1, 0x3, 0x5, 0x7, 0x9, 0xb, 0x2, 0xf, 0x11, 0x13, 0x17, 0x15, 0x19, 0xd, 0x18, 0x4, 0x8, 0x16, 0x6, 0x0, 0xa, 0xc, 0x14, 0x12, 0x10, 0xe},
	{0x4, 0x12, 0xe, 0x15, 0xf, 0x19, 0x9, 0x0, 0x18, 0x10, 0x14, 0x8, 0x11, 0x7, 0x17, 0xb, 0xd, 0x5, 0x13, 0x6, 0xa, 0x3, 0x2, 0xc, 0x16, 0x1},
	{0x15, 0x19, 0x1, 0x11, 0x6, 0x8, 0x13, 0x18, 0x14, 0xf, 0x12, 0x3, 0xd, 0x7, 0xb, 0x17, 0x0, 0x16, 0xc, 0x9, 0x10, 0xe, 0x5, 0x4, 0x2, 0xa},
	{0x9, 0xf, 0x6, 0x15, 0xe, 0x14, 0xc, 0x5, 0x18, 0x10, 0x1, 0x4, 0xd, 0x7, 0x19, 0x11, 0x3, 0xa, 0x0, 0x12, 0x17, 0xb, 0x8, 0x2, 0x13, 0x16},
	{0xd, 0x19, 0x9, 0x7, 0x6, 0x11, 0x2, 0x17, 0xc, 0x18, 0x12, 0x16, 0x1, 0xe, 0x14, 0x5, 0x0, 0x8, 0x15, 0xb, 0xf, 0x4, 0xa, 0x10, 0x3, 0x13},
	{0x5, 0xa, 0x10, 0x7, 0x13, 0xb, 0x17, 0xe, 0x2, 0x1, 0x9, 0x12, 0xf, 0x3, 0x19, 0x11, 0x0, 0xc, 0x4, 0x16, 0xd, 0x8, 0x14, 0x18, 0x6, 0x15},
}

func (r RotorNumber) ForwardWiring() []uint {
	return rotorForwardWiringValues[r]
}

func inverseWiring(wiring []uint) []uint {
	inverse := make([]uint, len(wiring))
	for i, f := range wiring {
		inverse[f] = uint(i)
	}
	return inverse
}

var rotorInverseWiringValues = [][]uint{
	{0x14, 0x16, 0x18, 0x6, 0x0, 0x3, 0x5, 0xf, 0x15, 0x19, 0x1, 0x4, 0x2, 0xa, 0xc, 0x13, 0x7, 0x17, 0x12, 0xb, 0x11, 0x8, 0xd, 0x10, 0xe, 0x9},
	{0x0, 0x9, 0xf, 0x2, 0x19, 0x16, 0x11, 0xb, 0x5, 0x1, 0x3, 0xa, 0xe, 0x13, 0x18, 0x14, 0x10, 0x6, 0x4, 0xd, 0x7, 0x17, 0xc, 0x8, 0x15, 0x12},
	{0x13, 0x0, 0x6, 0x1, 0xf, 0x2, 0x12, 0x3, 0x10, 0x4, 0x14, 0x5, 0x15, 0xd, 0x19, 0x7, 0x18, 0x8, 0x17, 0x9, 0x16, 0xb, 0x11, 0xa, 0xe, 0xc},
	{0x7, 0x19, 0x16, 0x15, 0x0, 0x11, 0x13, 0xd, 0xb, 0x6, 0x14, 0xf, 0x17, 0x10, 0x2, 0x4, 0x9, 0xc, 0x1, 0x12, 0xa, 0x3, 0x18, 0xe, 0x8, 0x5},
	{0x10, 0x2, 0x18, 0xb, 0x17, 0x16, 0x4, 0xd, 0x5, 0x13, 0x19, 0xe, 0x12, 0xc, 0x15, 0x9, 0x14, 0x3, 0xa, 0x6, 0x8, 0x0, 0x11, 0xf, 0x7, 0x1},
	{0x12, 0xa, 0x17, 0x10, 0xb, 0x7, 0x2, 0xd, 0x16, 0x0, 0x11, 0x15, 0x6, 0xc, 0x4, 0x1, 0x9, 0xf, 0x13, 0x18, 0x5, 0x3, 0x19, 0x14, 0x8, 0xe},
	{0x10, 0xc, 0x6, 0x18, 0x15, 0xf, 0x4, 0x3, 0x11, 0x2, 0x16, 0x13, 0x8, 0x0, 0xd, 0x14, 0x17, 0x5, 0xa, 0x19, 0xe, 0x12, 0xb, 0x7, 0x9, 0x1},
	{0x10, 0x9, 0x8, 0xd, 0x12, 0x0, 0x18, 0x3, 0x15, 0xa, 0x1, 0x5, 0x11, 0x14, 0x7, 0xc, 0x2, 0xf, 0xb, 0x4, 0x16, 0x19, 0x13, 0x6, 0x17, 0xe},
}

func (r RotorNumber) InverseWiring() []uint {
	return rotorInverseWiringValues[r]
}
