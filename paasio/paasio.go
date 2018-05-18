package paasio

import (
	"io"
	"sync"
)

type counter struct {
	n     int64
	nops  int
	rf    func(p []byte) (n int, err error)
	wf    func(p []byte) (n int, err error)
	mutex sync.Mutex
}

// ReadCount 1
func (c *counter) ReadCount() (n int64, nops int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.n, c.nops
}

// ReadCount 1
func (c *counter) WriteCount() (n int64, nops int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.n, c.nops
}

// Read 1
func (c *counter) Read(p []byte) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	n, err := c.rf(p)
	if err == nil {
		c.n += int64(n)
		c.nops++
	}
	return n, err
}

// Write 1
func (c *counter) Write(p []byte) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	n, err := c.wf(p)
	if err == nil {
		c.n += int64(n)
		c.nops++
	}
	return n, err
}

// NewReadCounter 1
func NewReadCounter(r io.Reader) ReadCounter {
	return &counter{n: 0, nops: 0, rf: r.Read, wf: nil}
}

// NewWriteCounter 1
func NewWriteCounter(w io.Writer) WriteCounter {
	return &counter{n: 0, nops: 0, rf: nil, wf: w.Write}
}

// NewReadWriteCounter 1
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &counter{n: 0, nops: 0, rf: rw.Read, wf: rw.Write}
}
