package main

import "time"

type Schedule interface {
	Next() time.Time
}
