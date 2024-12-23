package strings

import "testing"

func TestHasPrefix(t *testing.T) {
	if !HasPrefix("Hello, world", "Hello") {
		t.Error(`HasPrefix("Hello, world", "Hello") = false, want true`)
	}
}

func TestHasSuffix(t *testing.T) {
	if !HasSuffix("Hello, world", "world") {
		t.Error(`HasSuffix("Hello, world", "world") = false, want true`)
	}
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"été", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.s); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.s, got)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"été", true},
		{"A man, a plan, a canal: Panama.", true},
	}
	for _, test := range tests {
		if got := IsPalindrome2(test.s); got != test.want {
			t.Errorf("IsPalindrome2(%q) = %v", test.s, got)
		}
	}
}
