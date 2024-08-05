package infrastructure

import (
	"fmt"

	"github.com/beka-birhanu/grader/model"
)

// primaryKey represents a unique key for each subject based on name and ownerName.
type primaryKey struct {
	name      string
	ownerName string
}

// SubjectRepo is an in-memory repository for subjects.
type SubjectRepo struct {
	store map[primaryKey]*model.Subject
}

// New creates and returns a new SubjectRepo.
func New() *SubjectRepo {
	return &SubjectRepo{
		store: make(map[primaryKey]*model.Subject),
	}
}

// Add adds a new subject to the repository.
func (r *SubjectRepo) Add(subject *model.Subject) error {
	key := primaryKey{name: subject.Name(), ownerName: subject.OwnerName()}
	_, ok := r.store[key]
	if ok {
		return fmt.Errorf("subject name conflict, subject with name %s already exists", subject.Name())
	}

	r.store[key] = subject
	return nil
}

// ByOwner retrieves all subjects owned by the specified owner.
func (r *SubjectRepo) ByOwner(ownerName string) []*model.Subject {
	subjects := make([]*model.Subject, 0)
	for key, subject := range r.store {
		if key.ownerName == ownerName {
			subjects = append(subjects, subject)
		}
	}

	return subjects
}

