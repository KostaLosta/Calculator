// Package sqlstore contains a connection creator to sql database.
package sqlstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	// lib pq driver
	_ "github.com/lib/pq"
)

// Store - создаёт подключение к реляционной базе данных
type Store interface {
	DB() *sqlx.DB
	Close() (err error)
	Stats() sql.DBStats
}

type str struct {
	db *sqlx.DB
}

func (s *str) DB() *sqlx.DB {
	return s.db
}

func (s *str) Close() (err error) {
	err = s.db.Close()
	return
}

func (s *str) Stats() sql.DBStats {
	return s.db.Stats()
}

// New creates a private implementation of Store
func New(driverName string, dsn string, openConns int, lifeTime time.Duration) (s Store, err error) {
	dbConn, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		return
	}

	dbConn.SetConnMaxLifetime(lifeTime)
	dbConn.SetMaxOpenConns(openConns)

	s = &str{db: dbConn}
	return
}

// можно по пакетам разнести

type storeConn interface {
	DB() *sqlx.DB
}

// Storage интерфейс сервиса хранилища
type Storage interface {
	Create(name string, oneInt int, twoInt int, sum int) (err error)
}

type storage struct {
	storeConn
}

func (s *storage) Create(name string, oneInt int, twoInt int, sum int) (err error) {
	ctx := context.Background()

	// так делаются мептоды
	query := `INSERT INTO calculat (username, onenumber,twonumber, resultat)
			VALUES($1,$2, $3, $4)`

	_, err = s.DB().ExecContext(ctx, query, name, oneInt, twoInt, sum)
	if err != nil {
		return
	}
	return
}

// NewStorage создаёт экземпляр сервиса хранилища
func NewStorage(conn storeConn) (s Storage) {
	s = &storage{
		storeConn: conn,
	}
	return
}
