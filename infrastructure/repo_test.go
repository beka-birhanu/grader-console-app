package infrastructure

import (
	"testing"

	"github.com/beka-birhanu/grader/model"
)

func TestSubjectRepo_Add(t *testing.T) {
	repo := New()

	// Test adding a new subject
	subject1, err := model.New("Alice", "Math", 85)
	if err != nil {
		t.Fatalf("Failed to create subject: %v", err)
	}

	err = repo.Add(subject1)
	if err != nil {
		t.Errorf("Failed to add subject: %v", err)
	}

	// Test adding a duplicate subject
	err = repo.Add(subject1)
	if err == nil {
		t.Error("Expected error when adding a duplicate subject, got nil")
	}
}

func TestSubjectRepo_ByOwner(t *testing.T) {
	repo := New()

	// Test adding subjects
	subject1, err := model.New("Alice", "Math", 85)
	if err != nil {
		t.Fatalf("Failed to create subject: %v", err)
	}
	subject2, err := model.New("Alice", "Science", 90)
	if err != nil {
		t.Fatalf("Failed to create subject: %v", err)
	}
	subject3, err := model.New("Bob", "History", 75)
	if err != nil {
		t.Fatalf("Failed to create subject: %v", err)
	}

	_ = repo.Add(subject1)
	_ = repo.Add(subject2)
	_ = repo.Add(subject3)

	// Test retrieving subjects by owner
	subjectsAlice := repo.ByOwner("Alice")
	if len(subjectsAlice) != 2 {
		t.Errorf("Expected 2 subjects for Alice, got %d", len(subjectsAlice))
	}

	if subjectsAlice[0].Name() != "Math" || subjectsAlice[1].Name() != "Science" {
		t.Errorf("Subjects for Alice are not as expected")
	}

	subjectsBob := repo.ByOwner("Bob")
	if len(subjectsBob) != 1 {
		t.Errorf("Expected 1 subject for Bob, got %d", len(subjectsBob))
	}

	if subjectsBob[0].Name() != "History" {
		t.Errorf("Subject for Bob is not as expected")
	}
}

func TestSubjectRepo_ByOwner_NoSubjects(t *testing.T) {
	repo := New()

	// Test retrieving subjects by owner when no subjects are added
	subjects := repo.ByOwner("NonExistentOwner")
	if len(subjects) != 0 {
		t.Errorf("Expected 0 subjects for NonExistentOwner, got %d", len(subjects))
	}
}
