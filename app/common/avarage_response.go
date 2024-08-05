package common

import "github.com/beka-birhanu/grader/model"

// AvarageResponse holds the subjects and their calculated average grade.
type AvarageResponse struct {
	Subjects []*model.Subject // List of subjects
	Avarage  float32          // Calculated average grade
}

