package courses

// Store holds course data in memory.
type Store struct {
	Courses []Course
}

func NewStore() *Store {
	return &Store{
		Courses: []Course{
			{
				ID:          "course-1",
				Title:       "Introduction to Go",
				Category:    "Programming",
				Level:       "Beginner",
				Description: "Learn the basics of Go programming language.",
			},
			{
				ID:          "course-2",
				Title:       "Advanced Go Concurrency",
				Category:    "Programming",
				Level:       "Advanced",
				Description: "Deep dive into concurrency patterns in Go.",
			},
			{
				ID:          "course-3",
				Title:       "Web Development with Go",
				Category:    "Web Development",
				Level:       "Intermediate",
				Description: "Build web applications using Go and popular frameworks.",
			},
			{
				ID:          "course-4",
				Title:       "Data Structures and Algorithms in Go",
				Category:    "Computer Science",
				Level:       "Intermediate",
				Description: "Implement common data structures and algorithms in Go.",
			},
		},
	}
}

// GetAllCourses returns a snapshot of all courses in the store.
func (s *Store) GetAllCourses() []Course {
	return s.Courses
}
