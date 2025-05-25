package utils_test

import (
	"testing"

	"github.com/Myles-J/pokedexcli/internal/utils"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
		wantErr  bool
	}{
		{
			name:     "basic",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
			wantErr:  false,
		},
		{
			name:     "should return empty slice for empty string",
			input:    "",
			expected: []string{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.CleanInput(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("got nil; want error")
				}
			} else {
				if err != nil {
					t.Errorf("got error %v; want nil", err)
				}
			}

			if len(got) != len(tt.expected) {
				t.Errorf("got %v; want %v", got, tt.expected)
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("got %v; want %v", got, tt.expected)
					break
				}
			}
		})
	}
}
