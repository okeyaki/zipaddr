package flow

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestMerge(t *testing.T) {
	dataset := []struct {
		expected []byte
		line     string
	}{
		// 0ネスト
		{
			[]byte("0,1,2,3,4,A(BC)D,6,7,A(BC)D\n"),
			"0,1,2,3,4,A(B,6,7,A(B\n0,1,2,3,4,C)D,6,7,C)D\n",
		},
		// 1ネスト
		{
			[]byte("0,1,2,3,4,A(BCD)E,6,7,A(BCD)E\n"),
			"0,1,2,3,4,A(B,6,7,A(B\n0,1,2,3,4,C,6,7,C\n0,1,2,3,4,D)E,6,7,D)E\n",
		},
		// 2ネスト
		{
			[]byte("0,1,2,3,4,A(BC(DE)F)G,6,7,A(BC(DE)F)G\n"),
			"0,1,2,3,4,A(B,6,7,A(B\n0,1,2,3,4,C(D,6,7,C(D\n0,1,2,3,4,E)F)G,6,7,E)F)G\n",
		},
	}

	for _, data := range dataset {
		buf := bytes.NewBufferString(data.line)
		out := merge(buf)

		actual, err := ioutil.ReadAll(out)
		if err != nil {
			panic(err)
		}

		if bytes.Compare(data.expected, actual) != 0 {
			t.Errorf("\nexpected: %+v\nactual: %+v", string(data.expected), string(actual))
		}
	}
}
