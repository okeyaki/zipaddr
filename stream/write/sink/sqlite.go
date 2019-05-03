package sink

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"github.com/jmoiron/sqlx"
)

type SqliteSink struct {
	DataDir string

	db       *sqlx.DB
	revision string
}

func (s *SqliteSink) Open() error {
	db, err := sqlx.Open("sqlite3", s.DataDir+"/zipaddr.sqlite3.db")
	if err != nil {
		return err
	}
	s.db = db

	s.revision = time.Now().Format("20060102150405")

	if err := s.createRevisionTable(); err != nil {
		return err
	}

	if err := s.createDataTable(); err != nil {
		return err
	}

	if err := s.deleteOldTables(); err != nil {
		return err
	}

	return nil
}

func (s *SqliteSink) Close() error {
	return s.db.Close()
}

func (s *SqliteSink) Write(flow io.Reader) error {
	reader := csv.NewReader(flow)

	for {
		rec, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		if err := s.insertData(rec); err != nil {
			return err
		}
	}

	return nil
}

func (s *SqliteSink) createRevisionTable() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS revisions (revision VARCHAR(255))`)

	return err
}

func (s *SqliteSink) createDataTable() error {
	query := fmt.Sprintf(
		`CREATE TABLE data_%s (
				jiscode VARCHAR(255)
			, zipcode VARCHAR(255)
			, zipcode_old VARCHAR(255)
			, prefecture VARCHAR(255)
			, city VARCHAR(255)
			, town VARCHAR(255)
			, prefecture_ruby VARCHAR(255)
			, city_ruby VARCHAR(255)
			, town_ruby VARCHAR(255)
		)`,
		s.revision,
	)

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	if err := s.insertRevision(); err != nil {
		return err
	}

	return nil
}

func (s *SqliteSink) insertRevision() error {
	_, err := s.db.Exec(
		`INSERT INTO revisions (revision) VALUES (?)`,
		s.revision,
	)

	return err
}

func (s *SqliteSink) insertData(rec []string) error {
	query := fmt.Sprintf(
		`INSERT INTO data_%s (
				jiscode
			, zipcode_old
			, zipcode
			, prefecture_ruby
			, city_ruby
			, town_ruby
			, prefecture
			, city
			, town
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		s.revision,
	)

	_, err := s.db.Exec(
		query,
		rec[0],
		rec[1],
		rec[2],
		rec[3],
		rec[4],
		rec[5],
		rec[6],
		rec[7],
		rec[8],
	)

	return err
}

func (s *SqliteSink) deleteOldTables() error {
	tables, err := s.findOldRevisions()
	if err != nil {
		return err
	}

	for _, table := range tables {
		query := fmt.Sprintf("DROP TABLE IF EXISTS data_%s", table)

		_, err := s.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SqliteSink) findOldRevisions() ([]string, error) {
	rows, err := s.db.Queryx(
		`SELECT revision FROM revisions ORDER BY revision DESC LIMIT -1 OFFSET 12`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	revisions := []string{}
	for rows.Next() {
		var revision string
		if err := rows.Scan(&revision); err != nil {
			return nil, err
		}

		revisions = append(revisions, revision)
	}

	return revisions, nil
}
