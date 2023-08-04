package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	ID          string
	Name        string
	Description *string
	CategoryID  string
	db          *sql.DB
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryID string) (*Course, error) {
	id := uuid.New().String()
	query := "insert into courses(id,name,description,category_id) values ($1,$2,$3,$4)"

	_, err := c.db.Exec(query, id, name, description, categoryID)
	if err != nil {
		return &Course{}, err
	}

	return &Course{ID: id, Name: name, Description: &description, CategoryID: categoryID}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("Select id, description, name, category_id from courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		ct := Course{}
		if err := rows.Scan(&ct.ID, &ct.Description, &ct.Name, &ct.CategoryID); err != nil {
			return nil, err
		}
		courses = append(courses, ct)
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("Select id, description, name, category_id from courses where category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		ct := Course{}
		if err := rows.Scan(&ct.ID, &ct.Description, &ct.Name, &ct.CategoryID); err != nil {
			return nil, err
		}
		courses = append(courses, ct)
	}

	return courses, nil
}
