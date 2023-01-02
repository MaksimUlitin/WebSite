package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maksimUlitin/pkg/models/storage"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// главная страница
type neuteredFileSystem struct {
	fs http.FileSystem
}

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	pageBox  *storage.PageBoxModel
}

func main() {
	addr := flag.String("addr", ":8088", "Сетевой адрес HTTP")
	dsn := flag.String("dsn", "web:1_Qwertyuiop_2@/pageBox?parseTime=true", "Название MySQL источника данных")
	flag.Parse()

	// создание инфоЛогов
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	//передача данных (DSN) в openDB()
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	//Инициализируем новую структуру с зависимостями приложения.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		pageBox:  &storage.PageBoxModel{DB: db},
	}

	//новую структуру http.Server чтобы сервер использовал наш логгер
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("WEB-server запущен на порту %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// защита от проникновения в директорию к файлам
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {

	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
		}
	}
	return f, nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// ping() проверяет настройки подкулючения к бд
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
