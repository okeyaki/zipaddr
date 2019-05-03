package sink

import (
	"bufio"
	"fmt"
	"io"
)

type StreamSink struct {
	Stream io.Writer
}

func (s *StreamSink) Write(flow io.Reader) error {
	scanner := bufio.NewScanner(flow)
	for scanner.Scan() {
		fmt.Fprintln(s.Stream, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
