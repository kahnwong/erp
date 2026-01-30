package core

import (
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() string
		wantErr bool
	}{
		{
			name: "valid file exists",
			setup: func() string {
				tmpFile, err := os.CreateTemp("", "erp_test_*.txt")
				if err != nil {
					t.Fatalf("failed to create temp file: %v", err)
				}
				if _, err := tmpFile.WriteString("test content"); err != nil {
					t.Fatalf("failed to write to temp file: %v", err)
				}
				if err := tmpFile.Close(); err != nil {
					t.Fatalf("failed to close temp file: %v", err)
				}
				return tmpFile.Name()
			},
			wantErr: false,
		},
		{
			name: "file does not exist",
			setup: func() string {
				return "/tmp/nonexistent_file_12345.txt"
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := tt.setup()
			if !tt.wantErr {
				defer func() {
					_ = os.Remove(filename)
				}()
			}

			file, err := openFile(filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if file != nil {
				defer func() {
					_ = file.Close()
				}()
				if file.Name() != filename {
					t.Errorf("openFile() returned file with name %s, want %s", file.Name(), filename)
				}
			}
		})
	}
}
