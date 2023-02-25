package trace

import (
	"bytes"
	"io"
	"testing"
)

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("Error not be nil")
	} else {
		tracer.Trace("abc")
		if buf.String() != "abc\n" {
			t.Errorf("Trace should not write %s", buf.String())
		}
	}
}
