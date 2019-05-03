package flow

import (
	"io"
)

func Normalize(flow io.Reader) io.Reader {
	flow = Split(flow)

	return flow
}
