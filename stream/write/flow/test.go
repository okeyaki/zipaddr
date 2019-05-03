package flow

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func toShiftJIS(str string) []byte {
	reader := transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewEncoder())

	transformed, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return transformed
}
