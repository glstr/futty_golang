package context

import "testing"

func TestLogBufferPool(t *testing.T) {
	type Case struct {
		input  string
		expect string
	}

	cases := []Case{
		Case{
			input:  "hello world",
			expect: "hello world",
		},

		Case{
			input:  "debug",
			expect: "debug",
		},

		Case{
			input:  "smart dog",
			expect: "smart dog",
		},
	}

	for _, c := range cases {
		logBuf := getLogBuffer()
		logBuf.WriteLog(c.input)
		if logBuf.String() != c.expect {
			t.Errorf("expect:%s, real:%s", c.expect, logBuf.String())
		}
		putLogBuffer(logBuf)
	}
}
