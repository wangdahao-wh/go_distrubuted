package grades

import (
	"fmt"
	"sync"
)

type Studnet struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

func (s Studnet) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len((s.Grades)))
}

type GradeType string

type Students []Studnet

var (
	students      Students
	studentsMutex sync.Mutex
)

func (ss Students) GetByID(id int) (*Studnet, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}
	return nil, fmt.Errorf("student with id %d not found", id)
}

const (
	GradeQuiz = GradeType("quiz")
	GradeTest = GradeType("test")
	GradeExam = GradeType("exam")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
