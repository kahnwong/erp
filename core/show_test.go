package core

import (
	"os"
	"testing"
)

func TestReadData(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		wantCount   int
		wantErr     bool
	}{
		{
			name: "valid data with multiple items",
			fileContent: `food - milk - 2026-02-15 - 1
medicine - aspirin - 2026-03-20 - 2
household - soap - 2026-04-10 - 3`,
			wantCount: 3,
			wantErr:   false,
		},
		{
			name:        "empty file",
			fileContent: "",
			wantCount:   0,
			wantErr:     false,
		},
		{
			name: "file with malformed lines (skipped)",
			fileContent: `food - milk - 2026-02-15 - 1
invalid line
medicine - aspirin - 2026-03-20 - 2`,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:        "invalid quantity",
			fileContent: `food - milk - 2026-02-15 - abc`,
			wantCount:   0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temp file
			tmpFile, err := os.CreateTemp("", "erp_test_*.txt")
			if err != nil {
				t.Fatalf("failed to create temp file: %v", err)
			}
			defer func() {
				_ = os.Remove(tmpFile.Name())
			}()

			// Write test content
			if _, err := tmpFile.WriteString(tt.fileContent); err != nil {
				t.Fatalf("failed to write to temp file: %v", err)
			}
			if err := tmpFile.Close(); err != nil {
				t.Fatalf("failed to close temp file: %v", err)
			}

			// Setup config to use temp file
			originalConfig := AppConfig
			AppConfig = &Config{
				Path:       tmpFile.Name(),
				Categories: []string{"food", "medicine", "household"},
			}
			defer func() { AppConfig = originalConfig }()

			// Test readData
			items, err := readData()
			if (err != nil) != tt.wantErr {
				t.Errorf("readData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(items) != tt.wantCount {
				t.Errorf("readData() got %d items, want %d", len(items), tt.wantCount)
			}
		})
	}
}
