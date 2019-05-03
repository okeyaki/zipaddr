package flow

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestSplit(t *testing.T) {
	dataset := []struct {
		expected []byte
		line     string
	}{
		{
			[]byte("0,1,2,3,4,A,6,7,A\n"),
			"0,1,2,3,4,A,6,7,A\n",
		},
		{
			[]byte("0,1,2,3,4,A,6,7,A\n0,1,2,3,4,B,6,7,B\n"),
			"0,1,2,3,4,A、B,6,7,A、B\n",
		},
	}

	for _, data := range dataset {
		buf := bytes.NewBufferString(data.line)
		out := split(buf)

		actual, err := ioutil.ReadAll(out)
		if err != nil {
			panic(err)
		}

		if bytes.Compare(data.expected, actual) != 0 {
			t.Errorf("\nexpected: %+v\nactual: %+v", string(data.expected), string(actual))
		}
	}
}
