package database

import (
	"database/sql"

	"github.com/bikkusah/urlShortening/constant"
	"github.com/bikkusah/urlShortening/types"

	_ "github.com/mattn/go-sqlite3"
)

type manager struct {
	connection *sql.DB
}

var Mgr Manager

type Manager interface {
	Insert(url types.UrlDb) (int64, error)
	GetUrlFromCode(code string) (types.UrlDb, error)
}

func ConnectDb() {
	db, err := sql.Open("sqlite3", constant.Database)
	if err != nil {
		panic(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS url (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"url_code" TEXT NOT NULL UNIQUE,
		"long_url" TEXT,
		"short_url" TEXT,
		"created_at" INTEGER,
		"expired_at" INTEGER
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	Mgr = &manager{connection: db}
}

func (mgr *manager) Insert(url types.UrlDb) (int64, error) {
	insertSQL := `INSERT INTO url (url_code, long_url, short_url, created_at, expired_at) VALUES (?, ?, ?, ?, ?)`
	stmt, err := mgr.connection.Prepare(insertSQL)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(url.UrlCode, url.LongUrl, url.ShortUrl, url.CreatedAt, url.ExpiredAt)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (mgr *manager) GetUrlFromCode(code string) (types.UrlDb, error) {
	var url types.UrlDb

	querySQL := `SELECT url_code, long_url, short_url, created_at, expired_at FROM url WHERE url_code = ?`
	row := mgr.connection.QueryRow(querySQL, code)
	err := row.Scan(&url.UrlCode, &url.LongUrl, &url.ShortUrl, &url.CreatedAt, &url.ExpiredAt)
	if err != nil {
		return url, err
	}

	return url, nil
}
