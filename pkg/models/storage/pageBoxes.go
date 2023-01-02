package storage

import (
	"database/sql"
	"errors"
	"github.com/maksimUlitin/pkg/models"
)

type PageBoxModel struct {
	DB *sql.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (p *PageBoxModel) Insert(title, content, expires string) (int, error) {

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
func (p *PageBoxModel) Get(id int) (*models.PageBox, error) {
	//SQL запрос для получения данных одной записи
	stmt := `SELECT id,title,content, created, expires FROM pageBoxes
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := p.DB.QueryRow(stmt, id)

	s := &models.PageBox{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrorOnRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (p *PageBoxModel) Lastest() ([]*models.PageBox, error) {
	return nil, nil
}
