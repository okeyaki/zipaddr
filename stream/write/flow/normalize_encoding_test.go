package flow

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestNormalizeEncoding(t *testing.T) {
	dataset := []struct {
		expected []byte
		line     []byte
	}{
		{
			[]byte("あ"),
			toShiftJIS("あ"),
		},
	}

	for _, data := range dataset {
		buf := bytes.NewBuffer(data.line)
		out := normalizeEncoding(buf)

		actual, err := ioutil.ReadAll(out)
		if err != nil {
			panic(err)
		}

		if bytes.Compare(data.expected, actual) != 0 {
			t.Errorf("\nexpected: %+v\nactual: %+v", data.expected, actual)
		}
	}
}
