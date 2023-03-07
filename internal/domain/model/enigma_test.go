package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

var testEnigma = NewEnigma(
	&Rotor{
		number:      I,
		ringSetting: RingSetting(22),
		position:    RotorPosition(1),
	},
	&Rotor{
		number:      III,
		ringSetting: RingSetting(13),
		position:    RotorPosition(24),
	},
	&Rotor{
		number:      VI,
		ringSetting: RingSetting(5),
		position:    RotorPosition(12),
	},
	A,
	[]PlugboardCable{
		{from: 0, to: 25},
		{from: 13, to: 24},
		{from: 5, to: 12},
		{from: 14, to: 16},
		{from: 7, to: 22},
	},
)

func TestEnigma_String(t *testing.T) {
	tests := []struct {
		name  string
		value Enigma
		want  string
	}{
		{
			name:  "String",
			value: testEnigma,
			want:  "[I,22,1] [III,13,24] [VI,5,12] {A} (AZ,NY,FM,OQ,HW)",
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

func TestNewEnigmaFromText(t *testing.T) {
	testErr := errors.New("error parsing enigma values, " +
		"must define 3 rotors, one reflector and plugboard cables, like " +
		"[I,22,1] [III,13,24] [VI,5,12] {A} (AZ,YN,MF,OQ,WH)")
	tests := []struct {
		name    string
		value   string
		want    Enigma
		wantErr error
	}{
		{
			name:    "empty value",
			value:   "",
			want:    Enigma{},
			wantErr: testErr,
		},
		{
			name:    "random value",
			value:   "loren ipsum",
			want:    Enigma{},
			wantErr: testErr,
		},
		{
			name:    "wrong format",
			value:   "[I,22:1] [III,13,24] (AZ,YN,MF,OQ,WH)",
			want:    Enigma{},
			wantErr: testErr,
		},
		{
			name:    "missed reflector",
			value:   "[I,22:1] [III,13,24] [VI,5,12] (AZ,YN,MF,OQ,WH)",
			want:    Enigma{},
			wantErr: testErr,
		},
		{
			name:    "missed one rotor",
			value:   "[I,22,1] [III,13,24] {A} (AZ,YN,MF,OQ,WH)",
			want:    Enigma{},
			wantErr: testErr,
		},
		{
			name:    "wrong reflector",
			value:   "[I,22,1] [III,13,24] {D} (AZ,YN,MF,OQ,WH)",
			want:    Enigma{},
			wantErr: testErr,
		},
		{
			name:    "happy case",
			value:   "[I,22,1] [III,13,24] [VI,5,12] {A} (AZ,NY,FM,OQ,HW)",
			want:    testEnigma,
			wantErr: nil,
		},
		{
			name:  "no cables",
			value: "[I,22,1] [III,13,24] [VI,5,12] {B}",
			want: Enigma{
				leftRotor: &Rotor{
					number:      I,
					ringSetting: RingSetting(22),
					position:    RotorPosition(1),
				},
				middleRotor: &Rotor{
					number:      III,
					ringSetting: RingSetting(13),
					position:    RotorPosition(24),
				},
				rightRotor: &Rotor{
					number:      VI,
					ringSetting: RingSetting(5),
					position:    RotorPosition(12),
				},
				reflector: B,
				plugboard: Plugboard{
					cables: []PlugboardCable{},
					wiring: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
				},
			},
			wantErr: nil,
		},
		{
			name:    "error in roman number",
			value:   "[I,22,1] [III,13,24] [IIV,5,2] {C}",
			want:    Enigma{},
			wantErr: errors.New("invalid roman number"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEnigmaFromText(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestEnigma_Encrypt(t *testing.T) {
	text1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZAAAAAAAAAAAAAAAAAAAAAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBABCDEFGHIJKLMNOPQRSTUVWXYZ"
	text2 := strings.Repeat("VERYLONGTEXT", 100)
	tests := []struct {
		name   string
		enigma Enigma
		text   string
		want   string
	}{
		{
			"Case 1",
			getEnigma("[I,0,0] [II,0,0] [III,0,0] {B}"),
			text1,
			"BJELRQZVJWARXSNBXORSTNCFMEYHCXTGYJFLINHNXSHIUNTHEORXOPLOVFEKAGADSPNPCMHRVZCYECDAZIHVYGPITMSRZKGGHLSRBLHL",
		},
		{
			"VCase 2",
			getEnigma("[VII,1,10] [V,2,5] [IV,3,12] {B}"),
			text1,
			"FOTYBPKLBZQSGZBOPUFYPFUSETWKNQQHVNHLKJZZZKHUBEJLGVUNIOYSDTEZJQHHAOYYZSENTGXNJCHEDFHQUCGCGJBURNSEDZSEPLQP",
		},
		{
			"Case 3",
			getEnigma("[VII,11,11] [V,21,25] [IV,13,21] {A}"),
			text1,
			"NMZKCJSPLIACAUFRHEIIGEXJEYOOGNHLWEBCDOXQURFKKTSQBDOMVENAWCHSWHORCMVNGMLJJARHYGECJZNDKOTWDQISATHLFZRQDGIP",
		},
		{
			"Case 4",
			getEnigma("[VII,7,1] [V,17,5] [IV,11,11] {C}"),
			text1,
			"PCJXMPUBFYYCTPRDWYRWFBTLOXTTFXRLFSELGPHDCEFIGPQDRTERJCSUPHTUWAHOCDQKZVQYNRWTOXGPOPSQERGNDEBSTIXVTIGAYVUI",
		},
		{
			"Case 5",
			getEnigma("[III,11,3] [VI,13,5] [VIII,19,9] {B}"),
			text2,
			"PDGZFQPWNUOLTHLIJMJFGFVUFUAMKRTVGMAZGVVGEVBQCAZXETUQHIUKWCCYAQUIMCXWFYLSQQONAUMBDOYJNAWNFHECCL" +
				"WSRSCVCRWUCUOCPXPXAFTDZYNBURNIWZQFESLSYZYMGUVOMWJWPGGHEATIDPFVYVYWIURLAGFPUWBVYCCZJOYIMGZEVVDFR" +
				"HOPPWIOFJPJSMBDTWCBCXIZSVBGXQXBDCINLVQLDLDENGYAEBIAABZSFEWUMJNVTUSBHGMLFXUOGWFCRZYWBTKZXTKPVIKI" +
				"QLFLSNIVQJDLYPOYXCNEFMWMXXCXXRNWISRCMAVPZPUETHVHXOJQMRSSEWDSSDLKHPUPHJWPOFYDJMFECXTDDZTVHMOJSNM" +
				"FJCHAPBDLBVFSNCTASIBJMESFAYQYIHBCGZEJOLDVBIEKRADAYPKIDNSKJXPIBLSWSJVCSXUCGXCGNRAWIIKLHIZKKOVQVI" +
				"XVKOMWMKHHZWQTGXDLAIWMXAYCIVSAJKDFGXIQQUZMGRWKOVDNGFWFWYKCPKUQEHBDCFCNMWSTDFBCPCDNYMGIOLZNPIGHG" +
				"THXXTWDOUNLYBIOIODAWVNGUDALUPHRABCQXRTZLMFZTKPPWCZAMZUQABCZFUIFCTQLWTMVCMTZLFNUCLGRKBUTQKFDGRGU" +
				"JCCVWTQBOWUXEVDMWDIJNWQSNACUVJWWQVNWGHJEEZKKDUXLRSIDLQAKXYTOONTDHPHIWPOFSDBLSAUMUBVNWSWWFKEVHYQ" +
				"BVSQWEATFPVGMBAWWZKPSEQMCWPVLIWBOUXWCCIXQJUSGCLJESUDUGWDOHCVSOJWAGMPJTAESDUOFMNZGXSPHKMQZMLTKGS" +
				"FCXECWOAMWANZAMSJXMTTURNNTFFLLYYACRXUMHZEDHDEPLTFAKKEKXYKINPSETFRIUKNNEZGTIADHNFKZELQFBRKVBNCBQ" +
				"KCOHPBJXDZAGFYSZJCAVDOYPWMPFEWHZLLCSYVNBQKYPLTKODOJXTARJMJYORFOPXUCLPEZWMEWXKLHPNWGUAWGUMGPQSZQ" +
				"ZJNLYBHEQWAGFPXXTUCQZSZEDDMVQJDFZHGCQSTAXTJHNPBMLCYJPCARUZWZYSOWEQPMFZYKGTQLMJEZMMNPMJHTUSXXKBW" +
				"QMNDZIQSSIPAAEBBNOLICDYPNRYXVZGVUVSMARLGWZEIOKRTIXTPXHKTDNJOK",
		},
		{
			"Case 6",
			getEnigma("[I,0,0] [II,0,0] [III,0,0] {B} (AC,FG,JY,LW)"),
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			"QREBNMCYZELKQOJCGJVIVGLYEMUPCURPVPUMDIWXPPWROOQEGI",
		},
		{
			"Case 7",
			getEnigma("[IV,0,0] [VI,0,10] [III,0,6] {B} (BM,DH,RS,KN,GZ,FQ)"),
			"WRBHFRROSFHBCHVBENQFAGNYCGCRSTQYAJNROJAKVKXAHGUZHZVKWUTDGMBMSCYQSKABUGRVMIUOWAPKCMHYCRTSDEYTNJLVWNQY",
			"FYTIDQIBHDONUPAUVPNKILDHDJGCWFVMJUFNJSFYZTSPITBURMCJEEAMZAZIJMZAVFCTYTKYORHYDDSXHBLQWPJBMSSWIPSWLENZ",
		},
		{
			"Case 8",
			getEnigma("[I,5,0] [II,5,1] [III,4,20] {B} (AG,HR,YT,KI,FL,WE,NM,SD,OP,QJ)"),
			"RNXYAZUYTFNQFMBOLNYNYBUYPMWJUQSBYRHPOIRKQSIKBKEKEAJUNNVGUQDODVFQZHASHMQIHSQXICTSJNAUVZYIHVBBARPJADRH",
			"CFBJTPYXROYGGVTGBUTEBURBXNUZGGRALBNXIQHVBFWPLZQSCEZWTAWCKKPRSWOGNYXLCOTQAWDRRKBCADTKZGPWSTNYIJGLVIUQ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.enigma.Encrypt(tt.text)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestEnigma_Encrypt_and_Decrypt(t *testing.T) {
	t.Run("Encrypt and decrypt", func(t *testing.T) {
		enigma1 := getEnigma("[IV,0,0] [VI,0,10] [III,0,6] {B} (BM,DH,RS,KN,GZ,FQ)")
		clearText := strings.Repeat("TEXTTOENCRYPT", 50)
		encryptedText := enigma1.Encrypt(clearText)

		enigma2 := getEnigma("[IV,0,0] [VI,0,10] [III,0,6] {B} (BM,DH,RS,KN,GZ,FQ)")
		decryptedText := enigma2.Encrypt(encryptedText)

		assert.Equal(t, clearText, decryptedText)
	})

}

func getEnigma(text string) Enigma {
	enigma, err := NewEnigmaFromText(text)
	if err != nil {
		log.Fatal(err)
	}
	return enigma
}
