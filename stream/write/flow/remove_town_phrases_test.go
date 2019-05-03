package flow

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestRemoveTownPhrases(t *testing.T) {
	dataset := []struct {
		expected []byte
		line     string
	}{
		// 以下に掲載がない場合
		{
			[]byte("0,1,2,3,4,,6,7,\n"),
			"0,1,2,3,4,イカニケイサイガナイバアイ,6,7,以下に掲載がない場合\n",
		},
		// 地割
		{
			[]byte("0,1,2,3,4,A,6,7,A\n"),
			"0,1,2,3,4,A1チワリB,6,7,A1地割B\n",
		},
		{
			[]byte("0,1,2,3,4,A,6,7,A\n"),
			"0,1,2,3,4,Aダイ1チワリB,6,7,A第1地割B\n",
		},
		{
			[]byte("0,1,2,3,4,A,6,7,A\n"),
			"0,1,2,3,4,A(1チワリB,6,7,A(1地割B\n",
		},
		{
			[]byte("0,1,2,3,4,A,6,7,A\n"),
			"0,1,2,3,4,A(ダイ1チワリB,6,7,A(第1地割B\n",
		},
		// 小字
		{
			[]byte("0,1,2,3,4,AC,6,7,AC\n"),
			"0,1,2,3,4,A(B)C,6,7,A(B)C\n",
		},
	}

	for _, data := range dataset {
		buf := bytes.NewBufferString(data.line)
		out := removeTownPhrases(buf)

		actual, err := ioutil.ReadAll(out)
		if err != nil {
			panic(err)
		}

		if bytes.Compare(data.expected, actual) != 0 {
			t.Errorf("\nexpected: %+v\nactual: %+v", string(data.expected), string(actual))
		}
	}
}
