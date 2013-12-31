package main

import (
	"testing"
)

type UrlTestCase struct {
	url    string
	result bool
}

func TestIsUrl(t *testing.T) {
	tests := []UrlTestCase{
		UrlTestCase{"http://zerokspot.com", true},
		UrlTestCase{"https://zerokspot.com", true},
		UrlTestCase{"lala", false},
		UrlTestCase{"", false},
	}
	for _, test := range tests {
		actual := isUrl(test.url)
		if actual != test.result {
			t.Errorf("isUrl(%s) = %t (expected: %t)", test.url, actual, test.result)
		}
	}
}
