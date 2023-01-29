package main

import "testing"

func TestConvertToRoman(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{
			Description: "1 gets converted to I",
			Arabic:      1,
			Want:        "I",
		},
		{
			Description: "2 gets converted to II",
			Arabic:      2,
			Want:        "II",
		},
		{
			Description: "3 gets converted to III",
			Arabic:      3,
			Want:        "III",
		},
	}

	for _, tt := range cases {
		t.Run(tt.Description, func(t *testing.T) {
			got := ConvertToRoman(tt.Arabic)
			if got != tt.Want {
				t.Errorf("got %q, want %q", got, tt.Want)
			}

		})
	}
}
