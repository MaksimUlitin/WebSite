package storage

import (
	"database/sql"
	"github.com/maksimUlitin/pkg/models"
)

type PageBoxModel struct {
	DB *sql.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (p PageBoxModel) Insert(title, content, expires string) (int, error) {

	stmt := `INSERT INTO pageBoxes (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := p.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (p PageBoxModel) Get(id int) (*models.PageBox, error) {
	return nil, nil
}

// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (p PageBoxModel) Lastest() ([]*models.PageBox, error) {
	return nil, nil
}
