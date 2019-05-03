package flow

import (
	"encoding/csv"
	"io"
	"regexp"
	"strings"

	"github.com/tidwall/transform"
)

var (
	phrases = []string{
		"以下に掲載がない場合",
	}
	rubyPhrases = []string{
		"イカニケイサイガナイバアイ",
	}

	rePhrasePatterns = []*regexp.Regexp{
		regexp.MustCompile(`[\(]?(?:第)?[0-9]+地割.*`),
		regexp.MustCompile(`\(.+\)`),
	}
	reRubyPhrasePatterns = []*regexp.Regexp{
		regexp.MustCompile(`[\(]?(?:ダイ)?[0-9]+チワリ.*`),
		regexp.MustCompile(`\(.+\)`),
	}
)

func RemoveTownPhrases(flow io.Reader) io.Reader {
	flow = Merge(flow)
	flow = removeTownPhrases(flow)

	return flow
}

func removeTownPhrases(src io.Reader) io.Reader {
	reader := csv.NewReader(src)

	return transform.NewTransformer(func() ([]byte, error) {
		rec, err := reader.Read()
		if err != nil {
			return nil, err
		}

		for _, phrase := range phrases {
			rec[8] = strings.ReplaceAll(rec[8], phrase, "")
		}
		for _, phrase := range rubyPhrases {
			rec[5] = strings.ReplaceAll(rec[5], phrase, "")
		}

		for _, pattern := range rePhrasePatterns {
			rec[8] = pattern.ReplaceAllString(rec[8], "")
		}
		for _, pattern := range reRubyPhrasePatterns {
			rec[5] = pattern.ReplaceAllString(rec[5], "")
		}

		return []byte(strings.Join(rec, ",") + "\n"), nil
	})
}
