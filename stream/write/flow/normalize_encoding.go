package flow

import (
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func NormalizeEncoding(flow io.Reader) io.Reader {
	flow = normalizeEncoding(flow)

	return flow
}

func normalizeEncoding(src io.Reader) io.Reader {
	return transform.NewReader(src, japanese.ShiftJIS.NewDecoder())
}
