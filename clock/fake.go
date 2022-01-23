package clock

import "time"

type Fake struct {
	next time.Time
}

func NewFake() *Fake {
	return &Fake{}
}

func (f *Fake) SetNext(next time.Time) {
	f.next = next
}

func (f *Fake) Now() time.Time {
	return f.next
}
