package usecase_test

import "time"

type DummyClock struct {
	fixed time.Time
}

func (d *DummyClock) Now() time.Time {
	return d.fixed
}