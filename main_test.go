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
		{
			"4 gets converted to IV (can't repeat more than 3 times)",
			4,
			"IV",
		},
		{
			"5 gets converted to V",
			5,
			"V",
		},
		{
			"8 gets converted to VIII",
			8,
			"VIII",
		},
		{
			"9 gets converted to IX",
			9,
			"IX",
		},
		{
			"10 gets converted to X",
			10,
			"X",
		},
		{
			"14 gets converted to XIV",
			14,
			"XIV",
		},
		{
			"18 gets converted to XVIII",
			18,
			"XVIII",
		},
		{
			"20 gets converted to XX",
			20,
			"XX",
		},
		{
			"39 gets converted to XXXIX",
			39,
			"XXXIX",
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
