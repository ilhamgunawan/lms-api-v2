package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ilhamgunawan/lms-api-v2/db"
)

func CreateCourse(course db.Course) (c db.Course, err error) {
	course.ID = uuid.NewString()
	course.Status = "DRAFT"
	course.Created = time.Now().Format(time.RFC3339)

	err = db.GetDB().Insert(&course)

	if err != nil {
		return c, err
	}

	c = course

	return c, nil
}

func UpdateCourse(course db.Course) (c db.Course, err error) {
	query := ""

	if course.Title != "" {
		query += fmt.Sprintf("title = %s,", course.Title)
	}

	if course.Description != "" {
		query += fmt.Sprintf("description = '%s',", course.Description)
	}

	if course.Status != "" {
		query += fmt.Sprintf("status = '%s',", course.Status)
	}

	query += fmt.Sprintf("updated = '%s'", time.Now().Format(time.RFC3339))

	query = fmt.Sprintf("UPDATE course SET %s WHERE id='%s'", query, course.ID)

	_, err = db.GetDB().Exec(query)

	if err != nil {
		return c, err
	}

	c = course
	return c, nil
}

func DeleteCourse(courseId string) (err error) {
	deleted := time.Now().Format(time.RFC3339)
	query := fmt.Sprintf("UPDATE course SET deleted='%s' WHERE id='%s'", deleted, courseId)

	_, err = db.GetDB().Exec(query)

	if err != nil {
		return err
	}

	return nil
}
