package main

import (
	"fmt"
	"time"
)

type Fixed struct {
	duration time.Duration
	next     time.Time
}

func NewFixed(duration time.Duration) (*Fixed, error) {
	if duration < 0 {
		return nil, fmt.Errorf("invalid duration, must be >= 0")
	}
	return &Fixed{
		duration: duration,
		next:     time.Now().Add(duration),
	}, nil
}

func (f *Fixed) Next() time.Time {
	now := time.Now()
	if now.After(f.next) {
		f.next = f.next.Add(f.duration)
	}
	return f.next
}
