package main

import "testing"

func TestMasker(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test #1: проверка строки на содержание http://...",
			input:    "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			expected: "Here's my spammy page: http://******************* see you.",
		},
		{
			name:     "test #2: проверка строки на содержание http://...",
			input:    "http://yandex.com",
			expected: "http://**********",
		},
		{
			name:     "test #3: проверка строки на содержание http://...",
			input:    "http://yandex.com/maps , http://google.com",
			expected: "http://*************** , http://**********",
		},
		{
			name:     "test #4: проверка строки на содержание http://...",
			input:    "Some text",
			expected: "Some text",
		},
	}

	for _, test := range tests {
		result := masker(test.input)
		if result != test.expected {
			t.Errorf("Expected %s but got %s for input '%s'", test.expected, result, test.input)
		}
	}
}
