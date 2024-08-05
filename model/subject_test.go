package model

import (
	"testing"
)

// TestNewSubject tests the creation of new Subject instances.
func TestNewSubject(t *testing.T) {
	tests := []struct {
		name        string
		ownerName   string
		subjectName string
		mark        float32
		expectErr   bool
	}{
		{"Valid input", "John Doe", "Math", 85.5, false},
		{"Empty name", "John Doe", "", 85.5, true},
		{"Invalid mark (below range)", "John Doe", "Math", -1, true},
		{"Invalid mark (above range)", "John Doe", "Math", 101, true},
		{"Valid input with zero mark", "John Doe", "Math", 0, false},
		{"Valid input with perfect mark", "John Doe", "Math", 100, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subject, err := New(tt.ownerName, tt.subjectName, tt.mark)
			if (err != nil) != tt.expectErr {
				t.Errorf("New() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if err == nil {
				if subject.OwnerName() != tt.ownerName {
					t.Errorf("OwnerName() = %v, want %v", subject.OwnerName(), tt.ownerName)
				}
				if subject.Name() != tt.subjectName {
					t.Errorf("Name() = %v, want %v", subject.Name(), tt.subjectName)
				}
				if subject.Mark() != tt.mark {
					t.Errorf("Mark() = %v, want %v", subject.Mark(), tt.mark)
				}
			}
		})
	}
}

// TestValidateMark tests the validation of marks.
func TestValidateMark(t *testing.T) {
	tests := []struct {
		mark      float32
		expectErr bool
	}{
		{85.5, false},
		{-1, true},
		{101, true},
		{0, false},
		{100, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			err := validateMark(tt.mark)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateMark() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// TestValidateName tests the validation of subject names.
func TestValidateName(t *testing.T) {
	tests := []struct {
		name      string
		expectErr bool
	}{
		{"Math", false},
		{"", true},
		{"    ", true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			err := validateName(tt.name)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateName() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
