package sched

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

type Timer interface {
	Next() (next time.Time, done bool)
}

type Once struct {
	delay time.Duration
	done  bool
}

func NewOnce(delay time.Duration) (*Once, error) {
	if delay < 0 {
		return nil, fmt.Errorf("invalid delay, must be >= 0")
	}
	return &Once{
		delay: delay,
	}, nil
}

func NewOnceTime(t time.Time) (*Once, error) {
	remaining := time.Until(t)
	if remaining < 0 {
		return &Once{
			delay: remaining,
			done:  true,
		}, nil
	}
	return &Once{
		delay: remaining,
	}, nil
}

func (o *Once) Next() (time.Time, bool) {
	if !o.done {
		o.done = true
		return time.Now().Add(o.delay), false
	}
	return time.Time{}, o.done
}

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

func (f *Fixed) Next() (time.Time, bool) {
	now := time.Now()
	if now.After(f.next) {
		f.next = f.next.Add(f.duration)
	}
	return f.next, false
}

type Cron struct {
	expression cronexpr.Expression
}

func NewCron(cronExpression string) (*Cron, error) {
	expression, err := cronexpr.Parse(cronExpression)
	if err != nil {
		return nil, fmt.Errorf("cron expression invalid: %w", err)
	}
	return &Cron{expression: *expression}, nil
}

func (c *Cron) Next() (time.Time, bool) {
	return c.expression.Next(time.Now()), false
}
