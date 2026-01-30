package core

import (
	"testing"
)

func TestParseUserInput(t *testing.T) {
	// Setup test config
	AppConfig = &Config{
		Path:       "/tmp/test.txt",
		Categories: []string{"food", "medicine", "household"},
	}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "valid input with 3 args",
			args:    []string{"food", "milk", "2026-02-15"},
			wantErr: false,
		},
		{
			name:    "valid input with 4 args (quantity)",
			args:    []string{"food", "eggs", "2026-03-01", "12"},
			wantErr: false,
		},
		{
			name:    "invalid - too few args",
			args:    []string{"food", "milk"},
			wantErr: true,
		},
		{
			name:    "invalid - bad category",
			args:    []string{"invalid", "milk", "2026-02-15"},
			wantErr: true,
		},
		{
			name:    "invalid - bad date format",
			args:    []string{"food", "milk", "02-15-2026"},
			wantErr: true,
		},
		{
			name:    "invalid - bad quantity",
			args:    []string{"food", "milk", "2026-02-15", "abc"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := parseUserInput(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseUserInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if item.Category == "" || item.Item == "" || item.Date == "" {
					t.Errorf("parseUserInput() returned empty fields: %+v", item)
				}
				if item.Quantity < 1 {
					t.Errorf("parseUserInput() quantity should be at least 1, got %d", item.Quantity)
				}
			}
		})
	}
}
