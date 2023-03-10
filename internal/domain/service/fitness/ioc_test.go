package fitness

import "testing"

func Test_ioc_score(t *testing.T) {
	tests := []struct {
		name string
		text string
		want float32
	}{
		{
			"Case 1",
			"OZLUDYAKMGMXVFVARPMJIKVWPMBVWMOIDHYPLAYUWGBZFAFAFUQFZQISLEZMY",
			0.040437158197164536,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := ioc{}
			if got := i.score(tt.text); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}
