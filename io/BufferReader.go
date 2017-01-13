package buffer

import (
	"io"
	"sync"
	"bytes"
)

type BufferReader struct {
	mutex  sync.Mutex
	Buffer *bytes.Buffer
	reader io.Reader
	Error  error
	wait   sync.Cond
}

func NewBufferReader(r io.Reader, block ... bool) *BufferReader {
	reader := &BufferReader{
		Buffer:    &bytes.Buffer{},
		reader:    r,
	}
	reader.wait.L = &reader.mutex
	if block != nil && block[0] {
		defer func() {
			reader.drain()
		}()
	} else {
		go reader.drain()
	}
	return reader
}

func (r *BufferReader) drain() {
	buf := make([]byte, 1024)
	for {
		n, err := r.reader.Read(buf)
		r.mutex.Lock()
		if err != nil {
			r.Error = err
		} else {
			r.Buffer.Write(buf[0:n])
		}
		r.wait.Signal()
		r.mutex.Unlock()
		if err != nil {
			break
		}
	}
}

func (r *BufferReader) Read(p []byte) (n int, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for {
		n, err = r.Buffer.Read(p)
		if n > 0 {
			return n, err
		}
		if r.Error != nil {
			return 0, r.Error
		}
		r.wait.Wait()
	}
}

func (r *BufferReader) Close() error {
	closer, ok := r.reader.(io.ReadCloser)
	if !ok {
		return nil
	}
	return closer.Close()
}
