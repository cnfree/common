package common

import (
	"os"
	"io"
	"path/filepath"
	"sync"
	"bytes"
	"strings"
	"github.com/cnfree/common/io"
)

type ioUtil struct {
	mutex sync.Mutex
}

var IO = ioUtil{}

func (this ioUtil) CopyDir(source string, dest string) error {
	// get properties of source dir
	si, err := os.Stat(source)
	if err != nil {
		return err
	}
	err = os.MkdirAll(dest, si.Mode())
	if err != nil {
		return err
	}
	d, _ := os.Open(source)
	objects, err := d.Readdir(-1)
	for _, obj := range objects {
		sp := filepath.Join(source, "/", obj.Name())
		dp := filepath.Join(dest, "/", obj.Name())
		if obj.IsDir() {
			err = this.CopyDir(sp, dp)
			if err != nil {
				return err
			}
		} else {
			// perform copy
			err = this.CopyFile(sp, dp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this ioUtil) CopyFile(source string, dest string) error {
	ln, err := os.Readlink(source)
	if err == nil {
		return os.Symlink(ln, dest)
	}
	s, err := os.Open(source)
	if err != nil {
		return err
	}
	defer s.Close()
	d, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer d.Close()
	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}
	si, err := os.Stat(source)
	if err != nil {
		return err
	}
	err = os.Chmod(dest, si.Mode())
	return err
}

// WriteIndent writes indents to the buffer.
func (this ioUtil) WriteIndent(bf *bytes.Buffer, indent string, repead int) {
	bf.WriteString(strings.Repeat(indent, repead))
}

// UnifyLineFeed unifies line feeds.
func (this ioUtil) UnifyLineFeed(s string) string {
	return strings.Replace(strings.Replace(s, "\r\n", "\n", -1), "\r", "\n", -1)
}

//CaptureStdout not thread safe
func (this ioUtil) CaptureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func (this ioUtil) NewBufferReader(r io.Reader, block ... bool) *buffer.BufferReader {
	return buffer.NewBufferReader(r, block...)
}
