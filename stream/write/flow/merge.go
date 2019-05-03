package flow

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/tidwall/transform"
)

func Merge(flow io.Reader) io.Reader {
	flow = NormalizeCharacters(flow)
	flow = merge(flow)

	return flow
}

func merge(src io.Reader) io.Reader {
	reader := csv.NewReader(src)

	return transform.NewTransformer(func() ([]byte, error) {
		base, err := reader.Read()
		if err != nil {
			return nil, err
		}

		if 1 <= strings.Count(base[8], "(")-strings.Count(base[8], ")") {
			for {
				rec, err := reader.Read()
				if err != nil {
					return nil, err
				}

				base[8] += rec[8]
				base[5] += rec[5]

				if 0 == strings.Count(base[8], "(")-strings.Count(base[8], ")") {
					break
				}
			}
		}

		return []byte(strings.Join(base, ",") + "\n"), nil
	})
}
