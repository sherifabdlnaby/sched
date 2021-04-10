package main

import (
	"time"
)

type Timer interface {
	Next() time.Time
}
