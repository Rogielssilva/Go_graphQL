package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description *string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("insert into categories(id,name,description) values ($1,$2,$3)",
		id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: &description}, nil
}

func (c *Category) FindAll() ([]Category, error) {

	rows, err := c.db.Query("Select id, description, name from categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		ct := Category{}
		if err := rows.Scan(&ct.ID, &ct.Description, &ct.Name); err != nil {
			return nil, err
		}

		categories = append(categories, ct)
	}

	return categories, nil
}

func (c *Category) FindByCourseID(id string) (Category, error) {
	rows, err := c.db.Query("select c.id, c.description, c.name from categories c join courses co on c.id = co.category_id where co.id = $1", id)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()

	category := Category{}
	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Description, &category.Name); err != nil {
			return Category{}, err
		}
	}

	return category, nil
}
