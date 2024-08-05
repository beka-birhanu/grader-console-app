package gradesrvs

import (
	"github.com/beka-birhanu/grader/app/common"
	"github.com/beka-birhanu/grader/model"
)

// GradeService handles operations related to subjects and grades.
type GradeService struct {
	subjectRepo common.ISubjectRepo
}

// New creates a new instance of GradeService with the given subject repository.
func New(subjectRepo common.ISubjectRepo) *GradeService {
	return &GradeService{
		subjectRepo: subjectRepo,
	}
}

// Add creates a new subject and adds it to the repository.
// Returns the created subject and an error if any validation fails.
func (gs *GradeService) Add(ownerName, name string, mark float32) (*model.Subject, error) {
	subject, err := model.New(ownerName, name, mark)
	if err != nil {
		return subject, err
	}

	err = gs.subjectRepo.Add(subject)
	return subject, err
}

// Avarage calculates the average mark for all subjects owned by the given owner.
// Returns an AvarageResponse containing the subjects and the average mark.
func (gs *GradeService) Avarage(ownerName string) *common.AvarageResponse {
	var totalMark float32
	var average float32

	subjects := gs.subjectRepo.ByOwner(ownerName)

	for _, subject := range subjects {
		totalMark += subject.Mark()
	}

	if len(subjects) > 0 {
		average = totalMark / float32(len(subjects))
	}

	return &common.AvarageResponse{
		Subjects: subjects,
		Avarage:  average,
	}
}

