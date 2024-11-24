package grades

func init() {
	students = []Studnet{
		{
			ID:        1,
			FirstName: "NICK",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 78,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Roberto",
			LastName:  "Baggio",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 90,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 99,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Emma",
			LastName:  "Stone",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 77,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 80,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 78,
				},
			},
		},
		{
			ID:        4,
			FirstName: "Rachel",
			LastName:  "McAdams",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 98,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 99,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 94,
				},
			},
		},
		{
			ID:        5,
			FirstName: "Kelly",
			LastName:  "Clarkson",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 95,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 100,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 97,
				},
			},
		},
	}
}
