package trace

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Error("We haven't written our test yet")
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = off()
	silentTracer.Trace("something")
}
