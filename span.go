package tracer

import "time"

type status string

const (
	statusRunning status = "running"
	statusSuccess status = "success"
	statusFail    status = "fail"
	statusSkip    status = "skip"
)

type Span struct {
	id       string
	name     string
	duration time.Duration
	start    time.Time
	end      time.Time
	meta     any
	status   status
	reason   string
	children []*Span
	parent   *Span
}
