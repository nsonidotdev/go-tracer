package tracer

import (
	"time"

	"github.com/nsonidotdev/go-tracer/internal/id"
)

func Start(parent *Span, name string, meta map[string]string) (*Span, error) {
	id, err := id.GenerateRandomID(10)
	if err != nil {
		return nil, err
	}

	newSpan := &Span{
		id:     id,
		name:   name,
		start:  time.Now(),
		meta:   meta,
		status: statusRunning,
		parent: parent,
	}

	if parent != nil {
		parent.children = append(parent.children, newSpan)
	}

	return newSpan, nil
}
