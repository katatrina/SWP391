package main

import (
	"database/sql"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/katatrina/SWP391/internal/db/sqlc"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	infoLog        = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog       = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	dataSourceName = "postgres://postgres:12345@localhost:5432/bird_service_platform?sslmode=disable"
)

type application struct {
	infoLog        *log.Logger
	errorLog       *log.Logger
	db             *sql.DB
	store          *sqlc.Store
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	db, err := openDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Print("Connected to database")

	store := sqlc.NewStore(db)

	templateCache, err := initializeTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		infoLog:        infoLog,
		errorLog:       errorLog,
		db:             db,
		store:          store,
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	server := &http.Server{
		Addr:    "127.0.0.1:4000",
		Handler: app.routes(),
	}

	infoLog.Print("Starting server on http://localhost:4000")
	err = server.ListenAndServe()
	errorLog.Print(err)
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if pingErr := db.Ping(); pingErr != nil {
		return nil, err
	}

	return db, nil
}
