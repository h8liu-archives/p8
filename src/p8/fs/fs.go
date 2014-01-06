// package fs provides an simple file system wrapper
package fs

import (
	"io"
	"os"
)

func Open(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func Create(path string) (io.WriteCloser, error) {
	return os.Create(path)
}

func Exists(path string) (bool, error) {
	_, e := os.Stat(path)
	if os.IsNotExist(e) {
		return false, nil
	}
	if e != nil {
		return false, e
	}
	return true, nil
}
