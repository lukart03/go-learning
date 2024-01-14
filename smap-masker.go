package main

import "testing"

func masker(s string) string {
	var res []byte
	link := "http://"
	linkLen := len(link)
	linkFounnd := false

	for i := 0; i < len(s); i++ {
		if linkFounnd {
			if s[i] != ' ' {
				res = append(res, '*')
				continue
			}
			res = append(res, s[i])
			linkFounnd = false
			continue
		}
		if len(s[i:]) >= linkLen {
			if s[i:i+linkLen] == link {
				linkFounnd = true
				res = append(res, []byte(link)...)
				i += linkLen - 1
				continue
			}
		}

		res = append(res, s[i])

	}

	return string(res)
}

func testMasker(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Here's my spammy page: http://hehefouls.netHAHAHA see you.", "Here's my spammy page: http://******************* see you."},
		{"http://yandex.com/", "http://**********"},
		{"http://yandex.com/maps , http://google.com", "http://*************** , http://***********"},
		{"Some text", "Some text"},
	}

	for _, test := range tests {
		result := masker(test.input)
		if result != test.expected {
			t.Errorf("Expected %s but got %s for input '%s'", test.expected, result, test.input)
		}
	}
}
