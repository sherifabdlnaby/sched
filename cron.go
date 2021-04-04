package main

import (
	"github.com/gorhill/cronexpr"
	"time"
)

type Cron struct {
	expression cronexpr.Expression
}

func NewCron(cronExpression string) (*Cron, error) {
	expression, err := cronexpr.Parse(cronExpression)
	if err != nil {
		return nil, err
	}
	return &Cron{expression: *expression}, nil
}

func (c *Cron) Next() time.Time {
	return c.expression.Next(time.Now())
}
