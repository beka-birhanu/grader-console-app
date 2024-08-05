package gradesrvs

import (
	"testing"

	"github.com/beka-birhanu/grader/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for ISubjectRepo
type MockSubjectRepo struct {
	mock.Mock
}

func (m *MockSubjectRepo) Add(subject *model.Subject) error {
	args := m.Called(subject)
	return args.Error(0)
}

func (m *MockSubjectRepo) ByOwner(ownerName string) []*model.Subject {
	args := m.Called(ownerName)
	return args.Get(0).([]*model.Subject)
}

func TestAvarage(t *testing.T) {
	mockRepo := new(MockSubjectRepo)
	gs := New(mockRepo)

	// Define test cases
	tests := []struct {
		name            string
		ownerName       string
		subjects        []*model.Subject
		expectedAverage float32
	}{
		{"Multiple subjects", "John Doe", []*model.Subject{
			func() *model.Subject {
				subject, _ := model.New("John Doe", "Math", 85.5)
				return subject
			}(),
			func() *model.Subject {
				subject, _ := model.New("John Doe", "Science", 90.0)
				return subject
			}(),
		}, 87.75},
		{"Single subject", "John Doe", []*model.Subject{
			func() *model.Subject {
				subject, _ := model.New("John Doe", "Math", 75.0)
				return subject
			}(),
		}, 75.0},
		{"No subjects", "John Doe", []*model.Subject{}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("ByOwner", tt.ownerName).Return(tt.subjects).Once()

			response := gs.Avarage(tt.ownerName)
			assert.NotNil(t, response)
			assert.Equal(t, tt.expectedAverage, response.Avarage)

			mockRepo.AssertExpectations(t)
		})
	}
}

