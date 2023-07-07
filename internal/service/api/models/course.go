package models

import (
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/resources"
)

func newCourse(contract data.Contract) resources.Course {
	return resources.Course{
		Key: resources.NewKeyInt64(int64(contract.Id), resources.COURSE),
		Attributes: resources.CourseAttributes{
			Name:    contract.Name,
			Address: contract.Address,
		},
	}
}

func newCourseArray(contracts []data.Contract) []resources.Course {
	var courses = make([]resources.Course, 0)
	for _, contract := range contracts {
		courses = append(courses, newCourse(contract))
	}

	return courses
}

func NewCourseArrayResponse(contracts []data.Contract) resources.CourseListResponse {
	return resources.CourseListResponse{
		Data: newCourseArray(contracts),
	}
}
