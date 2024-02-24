package main

import "time"

type StructA struct {
	DateTime string
	Age      int
}

type StructB struct {
	// ab    time.Duration
	DateTime time.Time
	Age      int
}
