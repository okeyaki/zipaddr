package flow

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestNormalizeCharacters(t *testing.T) {
	dataset := []struct {
		expected []byte
		line     string
	}{
		// 英字
		{
			[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ\n"),
			"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ\n",
		},
		{
			[]byte("abcdefghijklmnopqrstuvwxyz\n"),
			"ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ\n",
		},
		// 数字
		{
			[]byte("123456789\n"),
			"１２３４５６７８９\n",
		},
		// カナ文字
		{
			[]byte("アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンヴガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポァィゥェォャュョッ\n"),
			"ｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜｦﾝｳﾞｶﾞｷﾞｸﾞｹﾞｺﾞｻﾞｼﾞｽﾞｾﾞｿﾞﾀﾞﾁﾞﾂﾞﾃﾞﾄﾞﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟｧｨｩｪｫｬｭｮｯ\n",
		},
		// 非 ASCII 記号
		{
			[]byte("ー\n"),
			"ー\n",
		},
		{
			[]byte("ー「 」、 。・\n"),
			"ー｢ ｣､ ｡･\n",
		},
		// ASCII 記号
		{
			[]byte("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~\n"),
			"！＂＃＄％＆＇（）＊＋，－．／：；＜＝＞？＠［＼］＾＿｀｛｜｝～\n",
		},
		// 空白文字
		{
			[]byte(" \n"),
			"　\n",
		},
	}

	for _, data := range dataset {
		buf := bytes.NewBufferString(data.line)
		out := normalizeCharacters(buf)

		actual, err := ioutil.ReadAll(out)
		if err != nil {
			panic(err)
		}

		if bytes.Compare(data.expected, actual) != 0 {
			t.Errorf("\nexpected: %+v\nactual: %+v", data.expected, actual)
		}
	}
}
