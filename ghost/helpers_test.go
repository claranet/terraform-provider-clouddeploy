package ghost

import (
	"reflect"
	"testing"
)

func TestStrToB64(t *testing.T) {
	cases := []struct {
		Input          string
		ExpectedOutput string
	}{
		{"mystring", "bXlzdHJpbmc="},
		{"", ""},
	}

	for _, tc := range cases {
		output := StrToB64(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from StrToB64.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestB64ToStr(t *testing.T) {
	cases := []struct {
		Input          string
		ExpectedOutput string
	}{
		{"bXlzdHJpbmc=", "mystring"},
		{"", ""},
		{"-1", ""},
		{"()", ""},
	}

	for _, tc := range cases {
		output := B64ToStr(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from B64ToStr.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestMatchesRegexp(t *testing.T) {
	cases := []struct {
		Function func(v interface{}, k string) (ws []string, errors []error)
		Value    string
		Valid    bool
	}{
		{MatchesRegexp(`^[a-zA-Z0-9]*$`), "thisIsAPositiveTest", true},
		{MatchesRegexp(`^[a-zA-Z0-9]*$`), "thisIsANegativeTest-", false},
	}

	for _, tc := range cases {
		_, err := tc.Function(tc.Value, tc.Value)
		if (tc.Valid && (err != nil)) || (!tc.Valid && (err == nil)) {
			t.Fatalf("Unexpected output from MatchesRegexp: %v", err)
		}
	}
}
