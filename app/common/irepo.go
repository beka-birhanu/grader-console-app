package common

import "github.com/beka-birhanu/grader/model"

// ISubjectRepo defines the interface for subject repository operations.
type ISubjectRepo interface {
	// Add adds a new subject to the repository.
	Add(subject *model.Subject) error

	// ByOwner retrieves all subjects owned by the specified owner.
	ByOwner(ownerName string) []*model.Subject
}

