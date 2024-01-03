package palindrome

import "testing"

func Test_isPalindromeNumber(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"test 1", 10, false},
		{"test 2", -121, false},
		{"test 3", 121, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeNumber(tt.input); got != tt.want {
				t.Errorf("isPalindromeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
