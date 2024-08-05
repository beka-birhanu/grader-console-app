package model

import (
	"fmt"
	"strings"
)

type Subject struct {
	ownerName string
	name      string
	mark      float32
}

// New creates new Subject and a returns a pointer or an error if
// any of the params fail on valdation
func New(ownerName, name string, mark float32) (*Subject, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := validateMark(mark); err != nil {
		return nil, err
	}

	return &Subject{ownerName: ownerName, name: name, mark: mark}, nil
}

func (s *Subject) Name() string {
	return s.name
}

func (s *Subject) OwnerName() string {
	return s.ownerName
}
func (s *Subject) Mark() float32 {
	return s.mark
}

// validateMark checks if the given mark is in range [0, 100].
// returns error if it is not.
func validateMark(mark float32) error {
	if mark < 0 || mark > 100 {
		return fmt.Errorf("grade must be in range [0, 100]")
	}
	return nil
}

// validateName checks if the given name is not empty.
// returns error if it is.
func validateName(name string) error {
	name = strings.Trim(name, " ")
	if len(name) == 0 {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}
