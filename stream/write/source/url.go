package source

import (
	"os"

	"github.com/okeyaki/zipaddr/util"
)

type URLSource struct {
	URL string

	file *os.File
}

func (s *URLSource) Open() error {
	paths, err := util.UnarchiveURL(s.URL)
	if err != nil {
		return err
	}

	file, err := os.Open(paths[0])
	if err != nil {
		return err
	}
	s.file = file

	return nil
}

func (s *URLSource) Close() error {
	return s.file.Close()
}

func (s URLSource) Read(b []byte) (int, error) {
	return s.file.Read(b)
}
