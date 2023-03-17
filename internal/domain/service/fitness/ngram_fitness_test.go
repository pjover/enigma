package fitness

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNgramFitness_Bigram_Score(t *testing.T) {
	tests := []struct {
		name             string
		ngramFitnessType NgramFitnessType
		text             string
		want             float32
		wantErr          error
	}{
		{
			"Bigram 1",
			Bigram,
			"OZLUDYAKMGMXVFVARPMJIKVWPMBVWMOIDHYPLAYUWGBZFAFAFUQFZQISLEZMY",
			-257.3492431640625,
			nil,
		},
		{
			"Bigram 2",
			Bigram,
			"PVBRDDLAGIHIFUJDFADORQOOMIZPYXDCBPWDSSNUSYZTJEWZPWFBWBMIEQXRF",
			-252.0023193359375,
			nil,
		},
		{
			"Bigram 3",
			Bigram,
			"ASZLOPPZRJKJSPPSTXKPUWYSKNMZZLHJDXJMMMDFODIHUBVCXMNICNYQBNQOD",
			-266.6638488769531,
			nil,
		},
		{
			"Bigram 4",
			Bigram,
			"FQLOGPZYXRJMTLMRKQAUQJPADHDZPFIKTQBFXAYMVSZPKXIQLOQCVRPKOBZSX",
			-293.7040100097656,
			nil,
		},
		{
			"Bigram 5",
			Bigram,
			"IUBAAJBRSNAFDMLLBVSYXISFXQZKQJRIQHOSHVYJXIFUZRMXWJVWHCCYHCXYG",
			-272.5676574707031,
			nil,
		},
		{
			"Bigram 6",
			Bigram,
			"RKMKBPWRDBXXRGABQBZRJDVHFPJZUSEBHWAEOGEUQFZEEBDCWNDHIAQDMHKPR",
			-266.79052734375,
			nil,
		},
		{
			"Bigram ",
			Bigram,
			"VYHQGRDYQIOEOLUBGBSNXWPZCHLDZQBWBEWOCQDBAFGUVHNGCIKXEIZGIZHPJ",
			-258.906982421875,
			nil,
		},
		{
			"Bigram 7",
			Bigram,
			"FCTMNNNAUXEVWTWACHOLOLSLTMDRZJZEVKKSSGUUTHVXXODSKTFGRUEIIXVWQ",
			-232.9383544921875,
			nil,
		},
		{
			"Bigram 8",
			Bigram,
			"YUIPIDBFPGLBYXZTCOQBCAHJYNSGDYLREYBRAKXGKQKWJEKWGAPTHGOMXJDSQ",
			-262.2741394042969,
			nil,
		},
		{
			"Bigram 9",
			Bigram,
			"KYHMFGOLXBSKVLGNZOAXGVTGXUIVFTGKPJU",
			-154.88406372070312,
			nil,
		},
		{
			"Bigram 10",
			Bigram,
			"OZLUDYAKMGMXVFVARPMJIKVWPMBVWMOIDHYPLAYUWGBZFAFAFUQFZQISLEZMYPVBRDDLAGIHIFUJDFADORQOOMIZPYXDCBPWDSSNUSYZTJEWZPWFBWBMIEQXRFASZLOPPZRJKJSPPSTXKPUWYSKNMZZLHJDXJMMMDFODIHUBVCXMNICNYQBNQODFQLOGPZYXRJMTLMRKQAUQJPADHDZPFIKTQBFXAYMVSZPKXIQLOQCVRPKOBZSXIUBAAJBRSNAFDMLLBVSYXISFXQZKQJRIQHOSHVYJXIFUZRMXWJVWHCCYHCXYGRKMKBPWRDBXXRGABQBZRJDVHFPJZUSEBHWAEOGEUQFZEEBDCWNDHIAQDMHKPRVYHQGRDYQIOEOLUBGBSNXWPZCHLDZQBWBEWOCQDBAFGUVHNGCIKXEIZGIZHPJFCTMNNNAUXEVWTWACHOLOLSLTMDRZJZEVKKSSGUUTHVXXODSKTFGRUEIIXVWQYUIPIDBFPGLBYXZTCOQBCAHJYNSGDYLREYBRAKXGKQKWJEKWGAPTHGOMXJDSQKYHMFGOLXBSKVLGNZOAXGVTGXUIVFTGKPJU",
			-2563.406982421875,
			nil,
		},
		{
			"Trigram 1",
			Trigram,
			"OZLUDYAKMGMXVFVARPMJIKVWPMBVWMOIDHYPLAYUWGBZFAFAFUQFZQISLEZMY",
			-443.9742736816406,
			nil,
		},
		{
			"Trigram 2",
			Trigram,
			"PVBRDDLAGIHIFUJDFADORQOOMIZPYXDCBPWDSSNUSYZTJEWZPWFBWBMIEQXRF",
			-427.6091613769531,
			nil,
		},
		{
			"Trigram 3",
			Trigram,
			"ASZLOPPZRJKJSPPSTXKPUWYSKNMZZLHJDXJMMMDFODIHUBVCXMNICNYQBNQOD",
			-455.9315490722656,
			nil,
		},
		{
			"Trigram 4",
			Trigram,
			"FQLOGPZYXRJMTLMRKQAUQJPADHDZPFIKTQBFXAYMVSZPKXIQLOQCVRPKOBZSX",
			-507.3169250488281,
			nil,
		},
		{
			"Trigram 5",
			Trigram,
			"IUBAAJBRSNAFDMLLBVSYXISFXQZKQJRIQHOSHVYJXIFUZRMXWJVWHCCYHCXYG",
			-461.4612121582031,
			nil,
		},
		{
			"Trigram 6",
			Trigram,
			"RKMKBPWRDBXXRGABQBZRJDVHFPJZUSEBHWAEOGEUQFZEEBDCWNDHIAQDMHKPR",
			-457.9267578125,
			nil,
		},
		{
			"Trigram 7",
			Trigram,
			"VYHQGRDYQIOEOLUBGBSNXWPZCHLDZQBWBEWOCQDBAFGUVHNGCIKXEIZGIZHPJ",
			-462.3972473144531,
			nil,
		},
		{
			"Trigram 8",
			Trigram,
			"FCTMNNNAUXEVWTWACHOLOLSLTMDRZJZEVKKSSGUUTHVXXODSKTFGRUEIIXVWQ",
			-408.73150634765625,
			nil,
		},
		{
			"Trigram 9",
			Trigram,
			"YUIPIDBFPGLBYXZTCOQBCAHJYNSGDYLREYBRAKXGKQKWJEKWGAPTHGOMXJDSQ",
			-442.8775939941406,
			nil,
		},
		{
			"Trigram 10",
			Trigram,
			"KYHMFGOLXBSKVLGNZOAXGVTGXUIVFTGKPJU",
			-276.1937561035156,
			nil,
		},
		{
			"Trigram 11",
			Trigram,
			"OZLUDYAKMGMXVFVARPMJIKVWPMBVWMOIDHYPLAYUWGBZFAFAFUQFZQISLEZMYPVBRDDLAGIHIFUJDFADORQOOMIZPYXDCBPWDSSNUSYZTJEWZPWFBWBMIEQXRFASZLOPPZRJKJSPPSTXKPUWYSKNMZZLHJDXJMMMDFODIHUBVCXMNICNYQBNQODFQLOGPZYXRJMTLMRKQAUQJPADHDZPFIKTQBFXAYMVSZPKXIQLOQCVRPKOBZSXIUBAAJBRSNAFDMLLBVSYXISFXQZKQJRIQHOSHVYJXIFUZRMXWJVWHCCYHCXYGRKMKBPWRDBXXRGABQBZRJDVHFPJZUSEBHWAEOGEUQFZEEBDCWNDHIAQDMHKPRVYHQGRDYQIOEOLUBGBSNXWPZCHLDZQBWBEWOCQDBAFGUVHNGCIKXEIZGIZHPJFCTMNNNAUXEVWTWACHOLOLSLTMDRZJZEVKKSSGUUTHVXXODSKTFGRUEIIXVWQYUIPIDBFPGLBYXZTCOQBCAHJYNSGDYLREYBRAKXGKQKWJEKWGAPTHGOMXJDSQKYHMFGOLXBSKVLGNZOAXGVTGXUIVFTGKPJU",
			-4490.65869140625,
			nil,
		},
		{
			"Quadram 1",
			Quadram,
			"OZLUDYAKMGMXVFVARPMJIKVWPMBVWMOIDHYPLAYUWGBZFAFAFUQFZQISLEZMY",
			-523.5413208007812,
			nil,
		},
		{
			"Quadram 2",
			Quadram,
			"PVBRDDLAGIHIFUJDFADORQOOMIZPYXDCBPWDSSNUSYZTJEWZPWFBWBMIEQXRF",
			-531.3290405273438,
			nil,
		},
		{
			"Quadram 3",
			Quadram,
			"ASZLOPPZRJKJSPPSTXKPUWYSKNMZZLHJDXJMMMDFODIHUBVCXMNICNYQBNQOD",
			-539.1807861328125,
			nil,
		},
		{
			"Quadram 4",
			Quadram,
			"FQLOGPZYXRJMTLMRKQAUQJPADHDZPFIKTQBFXAYMVSZPKXIQLOQCVRPKOBZSX",
			-546.4402465820312,
			nil,
		},
		{
			"Quadram 5",
			Quadram,
			"IUBAAJBRSNAFDMLLBVSYXISFXQZKQJRIQHOSHVYJXIFUZRMXWJVWHCCYHCXYG",
			-542.9093017578125,
			nil,
		},
		{
			"Quadram 6",
			Quadram,
			"RKMKBPWRDBXXRGABQBZRJDVHFPJZUSEBHWAEOGEUQFZEEBDCWNDHIAQDMHKPR",
			-535.1298217773438,
			nil,
		},
		{
			"Quadram 7",
			Quadram,
			"VYHQGRDYQIOEOLUBGBSNXWPZCHLDZQBWBEWOCQDBAFGUVHNGCIKXEIZGIZHPJ",
			-544.1270751953125,
			nil,
		},
		{
			"Quadram 8",
			Quadram,
			"FCTMNNNAUXEVWTWACHOLOLSLTMDRZJZEVKKSSGUUTHVXXODSKTFGRUEIIXVWQ",
			-509.68878173828125,
			nil,
		},
		{
			"Quadram 9",
			Quadram,
			"YUIPIDBFPGLBYXZTCOQBCAHJYNSGDYLREYBRAKXGKQKWJEKWGAPTHGOMXJDSQ",
			-533.2517700195312,
			nil,
		},
		{
			"Quadram 10",
			Quadram,
			"KYHMFGOLXBSKVLGNZOAXGVTGXUIVFTGKPJU",
			-301.8176574707031,
			nil,
		},
		{
			"Quadram 11",
			Quadram,
			"OZLUDYAKMGMXVFVARPMJIKVWPMBVWMOIDHYPLAYUWGBZFAFAFUQFZQISLEZMYPVBRDDLAGIHIFUJDFADORQOOMIZPYXDCBPWDSSNUSYZTJEWZPWFBWBMIEQXRFASZLOPPZRJKJSPPSTXKPUWYSKNMZZLHJDXJMMMDFODIHUBVCXMNICNYQBNQODFQLOGPZYXRJMTLMRKQAUQJPADHDZPFIKTQBFXAYMVSZPKXIQLOQCVRPKOBZSXIUBAAJBRSNAFDMLLBVSYXISFXQZKQJRIQHOSHVYJXIFUZRMXWJVWHCCYHCXYGRKMKBPWRDBXXRGABQBZRJDVHFPJZUSEBHWAEOGEUQFZEEBDCWNDHIAQDMHKPRVYHQGRDYQIOEOLUBGBSNXWPZCHLDZQBWBEWOCQDBAFGUVHNGCIKXEIZGIZHPJFCTMNNNAUXEVWTWACHOLOLSLTMDRZJZEVKKSSGUUTHVXXODSKTFGRUEIIXVWQYUIPIDBFPGLBYXZTCOQBCAHJYNSGDYLREYBRAKXGKQKWJEKWGAPTHGOMXJDSQKYHMFGOLXBSKVLGNZOAXGVTGXUIVFTGKPJU",
			-5361.82470703125,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut, err := NewNgramFitness(tt.ngramFitnessType)
			actual := sut.Score(tt.text)
			assert.Equal(t, tt.want, actual)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}