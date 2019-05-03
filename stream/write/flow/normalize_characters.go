package flow

import (
	"bufio"
	"io"

	"github.com/tidwall/transform"
	"golang.org/x/text/unicode/norm"
)

func NormalizeCharacters(flow io.Reader) io.Reader {
	flow = NormalizeEncoding(flow)
	flow = normalizeCharacters(flow)

	return flow
}

func normalizeCharacters(src io.Reader) io.Reader {
	reader := bufio.NewReader(src)

	return transform.NewTransformer(func() ([]byte, error) {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		return norm.NFKC.Bytes([]byte(line)), nil
	})
}
