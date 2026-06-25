package main

import (
	"reflect"
	"testing"
)

func TestRenderLines(t *testing.T) {
	banner := map[rune][]string{
		'A': {"A0", "A1", "A2", "A3", "A4", "A5", "A6", "A7"},
		'B': {"B0", "B1", "B2", "B3", "B4", "B5", "B6", "B7"},
		'C': {"C0", "C1", "C2", "C3", "C4", "C5", "C6", "C7"},
		' ': {"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
	}

	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "1 - single character A",
			input: "A",
			want:  []string{"A0", "A1", "A2", "A3", "A4", "A5", "A6", "A7"},
		},
		{
			name:  "2 - single character B",
			input: "B",
			want:  []string{"B0", "B1", "B2", "B3", "B4", "B5", "B6", "B7"},
		},
		{
			name:  "3 - two characters AB",
			input: "AB",
			want:  []string{"A0B0", "A1B1", "A2B2", "A3B3", "A4B4", "A5B5", "A6B6", "A7B7"},
		},
		{
			name:  "4 - two characters BA",
			input: "BA",
			want:  []string{"B0A0", "B1A1", "B2A2", "B3A3", "B4A4", "B5A5", "B6A6", "B7A7"},
		},
		{
			name:  "5 - three characters ABC",
			input: "ABC",
			want:  []string{"A0B0C0", "A1B1C1", "A2B2C2", "A3B3C3", "A4B4C4", "A5B5C5", "A6B6C6", "A7B7C7"},
		},
		{
			name:  "6 - repeated characters AAA",
			input: "AAA",
			want:  []string{"A0A0A0", "A1A1A1", "A2A2A2", "A3A3A3", "A4A4A4", "A5A5A5", "A6A6A6", "A7A7A7"},
		},
		{
			name:  "7 - repeated characters BBB",
			input: "BBB",
			want:  []string{"B0B0B0", "B1B1B1", "B2B2B2", "B3B3B3", "B4B4B4", "B5B5B5", "B6B6B6", "B7B7B7"},
		},
		{
			name:  "8 - characters with space A B",
			input: "A B",
			want:  []string{"A0  B0", "A1  B1", "A2  B2", "A3  B3", "A4  B4", "A5  B5", "A6  B6", "A7  B7"},
		},
		{
			name:  "9 - starts with space",
			input: " AB",
			want:  []string{"  A0B0", "  A1B1", "  A2B2", "  A3B3", "  A4B4", "  A5B5", "  A6B6", "  A7B7"},
		},
		{
			name:  "10 - ends with space",
			input: "AB ",
			want:  []string{"A0B0  ", "A1B1  ", "A2B2  ", "A3B3  ", "A4B4  ", "A5B5  ", "A6B6  ", "A7B7  "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RenderLines(tt.input, banner)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"RenderLines(%q) = %v, want %v",
					tt.input,
					got,
					tt.want,
				)
			}
		})
	}
}
