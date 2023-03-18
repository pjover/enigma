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

var rotorStringValues = []string{
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
	return rotorStringValues[r]
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

func decodeWiring(encoding string) []int {
	charWiring := []rune(encoding)
	wiring := make([]int, len(charWiring))
	for i, c := range charWiring {
		wiring[i] = int(c - 'A')
	}
	return wiring
}

var rotorForwardWiringValues = [][]int{
	{4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9},
	{0, 9, 3, 10, 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4},
	{1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23, 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14},
	{4, 18, 14, 21, 15, 25, 9, 0, 24, 16, 20, 8, 17, 7, 23, 11, 13, 5, 19, 6, 10, 3, 2, 12, 22, 1},
	{21, 25, 1, 17, 6, 8, 19, 24, 20, 15, 18, 3, 13, 7, 11, 23, 0, 22, 12, 9, 16, 14, 5, 4, 2, 10},
	{9, 15, 6, 21, 14, 20, 12, 5, 24, 16, 1, 4, 13, 7, 25, 17, 3, 10, 0, 18, 23, 11, 8, 2, 19, 22},
	{13, 25, 9, 7, 6, 17, 2, 23, 12, 24, 18, 22, 1, 14, 20, 5, 0, 8, 21, 11, 15, 4, 10, 16, 3, 19},
	{5, 10, 16, 7, 19, 11, 23, 14, 2, 1, 9, 18, 15, 3, 25, 17, 0, 12, 4, 22, 13, 8, 20, 24, 6, 21},
}

func (r RotorNumber) ForwardWiring() []int {
	return rotorForwardWiringValues[r]
}

func inverseWiring(wiring []int) []int {
	inverse := make([]int, len(wiring))
	for i, f := range wiring {
		inverse[f] = i
	}
	return inverse
}

var rotorInverseWiringValues = [][]int{
	{20, 22, 24, 6, 0, 3, 5, 15, 21, 25, 1, 4, 2, 10, 12, 19, 7, 23, 18, 11, 17, 8, 13, 16, 14, 9},
	{0, 9, 15, 2, 25, 22, 17, 11, 5, 1, 3, 10, 14, 19, 24, 20, 16, 6, 4, 13, 7, 23, 12, 8, 21, 18},
	{19, 0, 6, 1, 15, 2, 18, 3, 16, 4, 20, 5, 21, 13, 25, 7, 24, 8, 23, 9, 22, 11, 17, 10, 14, 12},
	{7, 25, 22, 21, 0, 17, 19, 13, 11, 6, 20, 15, 23, 16, 2, 4, 9, 12, 1, 18, 10, 3, 24, 14, 8, 5},
	{16, 2, 24, 11, 23, 22, 4, 13, 5, 19, 25, 14, 18, 12, 21, 9, 20, 3, 10, 6, 8, 0, 17, 15, 7, 1},
	{18, 10, 23, 16, 11, 7, 2, 13, 22, 0, 17, 21, 6, 12, 4, 1, 9, 15, 19, 24, 5, 3, 25, 20, 8, 14},
	{16, 12, 6, 24, 21, 15, 4, 3, 17, 2, 22, 19, 8, 0, 13, 20, 23, 5, 10, 25, 14, 18, 11, 7, 9, 1},
	{16, 9, 8, 13, 18, 0, 24, 3, 21, 10, 1, 5, 17, 20, 7, 12, 2, 15, 11, 4, 22, 25, 19, 6, 23, 14},
}

func (r RotorNumber) InverseWiring() []int {
	return rotorInverseWiringValues[r]
}

func (r RotorNumber) IsAtNotch(currentPosition RotorPosition) bool {
	switch r {
	case I:
		return currentPosition == 16
	case II:
		return currentPosition == 4
	case III:
		return currentPosition == 21
	case IV:
		return currentPosition == 9
	case V:
		return currentPosition == 25
	default:
		return currentPosition == 12 || currentPosition == 25
	}
}
