package sequence

import "time"

type epoch struct {
	currMillis int64
	lastMillis int64
}

func (t *epoch) next() *epoch {
	t.currMillis = int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
	return t
}

func (t *epoch) get() int64 {
	return t.currMillis
}

func (t *epoch) toCopy() {
	t.lastMillis = t.get()
}

func (t *epoch) isEqual() bool {
	return t.get() == t.lastMillis
}

func (t *epoch) waitNotEqual() {
	for t.isEqual() {
		t.next()
	}
}
