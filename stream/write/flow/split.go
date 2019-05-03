package flow

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/tidwall/transform"
)

func Split(flow io.Reader) io.Reader {
	flow = RemoveTownPhrases(flow)
	flow = split(flow)

	return flow
}

func split(src io.Reader) io.Reader {
	reader := csv.NewReader(src)

	return transform.NewTransformer(func() ([]byte, error) {
		base, err := reader.Read()
		if err != nil {
			return nil, err
		}

		towns := strings.Split(base[8], "、")
		townRubys := strings.Split(base[5], "、")

		recs := [][]string{}
		for i := range towns {
			rec := make([]string, len(base))
			copy(rec, base)

			rec[8] = towns[i]
			rec[5] = townRubys[i]

			recs = append(recs, rec)
		}

		joined := ""
		for _, rec := range recs {
			joined += strings.Join(rec, ",") + "\n"
		}

		return []byte(joined), nil
	})
}
