package fortesting

import "testing"

func TestError(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"No Error", "No Error"},
	}

	for _, test := range tests {
		got := Error(test.input)
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}
