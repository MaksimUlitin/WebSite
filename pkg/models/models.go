package models

import (
	"errors"
	"time"
)

var ErrorOnRecord = errors.New("models: запрашиваемая запить не найдена")

type PageBox struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
