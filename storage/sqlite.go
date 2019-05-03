package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SqliteStorage struct {
	DataDir string

	db       *sqlx.DB
	revision string
}

func (s *SqliteStorage) Open() error {
	db, err := sqlx.Open("sqlite3", s.DataDir+"/zipaddr.sqlite3.db")
	if err != nil {
		return err
	}
	s.db = db

	revision, err := s.findRevision()
	if err != nil {
		return err
	}
	s.revision = revision

	return nil
}

func (s *SqliteStorage) Close() error {
	return s.db.Close()
}

func (s *SqliteStorage) FindByZipcode(zipcode string) ([]Address, error) {
	query := fmt.Sprintf(
		`SELECT * FROM data_%s WHERE zipcode = ? ORDER BY zipcode`,
		s.revision,
	)

	rows, err := s.db.Queryx(query, zipcode)
	if err != nil {
		return nil, err
	}

	addrs := []Address{}
	for rows.Next() {
		addr := Address{}

		if err := rows.StructScan(&addr); err != nil {
			return nil, err
		}

		addrs = append(addrs, addr)
	}

	return addrs, nil
}

func (s *SqliteStorage) findRevision() (string, error) {
	rows, err := s.db.Queryx(`SELECT revision FROM revisions ORDER BY revision DESC`)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var revision string
		if err := rows.Scan(&revision); err != nil {
			return "", err
		}
		return revision, nil
	}

	return "", fmt.Errorf("TBD")
}
