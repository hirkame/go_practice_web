package trace

import (
	"fmt"
	"io"
	"testing"
)

// Tracer is an interface that can record events in this program.
type Tracer interface {
	// 任意の型を何個でも受け取れる。
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

// New tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// TestOff Tracer
func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("Data")
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off Tracer function and return empty tracer.
func Off() Tracer {
	return &nilTracer{}
}
