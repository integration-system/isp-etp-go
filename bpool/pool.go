package bpool

import (
	"bytes"
	"sync"
)

var bpool sync.Pool

func Get() *bytes.Buffer {
	b, ok := bpool.Get().(*bytes.Buffer)
	if !ok {
		b = bytes.NewBuffer(make([]byte, 0, 4096))
	}
	return b
}

func Put(b *bytes.Buffer) {
	b.Reset()
	bpool.Put(b)
}
