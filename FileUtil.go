package common

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"runtime"
	"bufio"
	"crypto/md5"
	"io"
)

const slash = "/"

type fileUtil struct {
	mutex sync.Mutex
}

var File = fileUtil{}

// Exists returns a boolean indicating whether the specified file path already exists.
func (this fileUtil) Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}
	return true
}

// IsFile returns a boolean indicating whether the specified file path is a file.
//
// If there is an error, it will be of type *PathError.
func (this fileUtil) IsFile(filePath string) (bool, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		return true, nil
	default:
		return false, nil
	}
}

// IsDirectory returns a boolean indicating whether the specified file path is a directory.
//
// If there is an error, it will be of type *PathError.
func (this fileUtil) IsDirectory(filePath string) (bool, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true, nil
	default:
		return false, nil
	}
}

// MakeDir creates a new directory with the specified file path and permission bits.
//
// If there is an error, it will be of type *PathError.
func (this fileUtil) MakeDir(filePath string, perm os.FileMode) (bool, error) {
	err := os.Mkdir(filePath, perm)
	return err == nil, err
}

// MakeRootDir creates the named directory with mode 0777 (before umask)
//
// If there is an error, it will be of type *PathError.
func (this fileUtil) MakeRootDir(filePath string) (bool, error) {
	result, err := this.MakeDir(filePath, os.ModePerm)
	return result, err
}

// MakeDirs creates a directory named path, along with any necessary parents.
// The permission bits perm are used for all directories that MakeDirs creates.
// If path is already a directory, MakeDirs does nothing.
//
// If there is an error, it will be of type *PathError.
func (this fileUtil) MakeDirs(filePath string, perm os.FileMode) (bool, error) {
	err := os.MkdirAll(filePath, perm)
	return err == nil, err
}

// MakeRootDirs creates a directory named path, along with any necessary parents.
// Mode 0777 are used for all directories that MakeRootDirs creates.

// If there is an error, it will be of type *PathError.
func (this fileUtil) MakeRootDirs(filePath string) (bool, error) {
	result, err := this.MakeDirs(filePath, os.ModePerm)
	return result, err
}

// Length in bytes for regular files or length count for directory.
func (this fileUtil) Length(filePath string) (int64, error) {
	result, err := this.IsFile(filePath)
	if err != nil {
		return 0, err
	}
	if result {
		fi, err := os.Stat(filePath)
		if err != nil {
			return 0, err
		}
		return fi.Size(), nil
	} else {
		children, err := this.List(filePath)
		if err != nil {
			return 0, err
		}
		var totalLength int64 = 0
		for _, child := range children {
			length, _ := this.Length(filepath.Join(filePath, child))
			totalLength += length
		}
		return totalLength, nil
	}
}

// List reads the specified file path and returns
// a list of directory or file entries sorted by file name.
func (this fileUtil) List(filePath string) ([]string, error) {
	if result, err := this.IsFile(filePath); err != nil || result {
		return nil, err
	}

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return nil, err
	}

	children := make([]string, len(files))
	for index, f := range files {
		children[index] = f.Name()
	}
	return children, nil
}

// Delete removes the named file or directory.
//
// If there is an error, it will be of type *PathError.
func (this fileUtil) Delete(filePath string) (bool, error) {
	err := os.Remove(filePath)
	return err == nil, err
}

// DeleteAll removes path and any children it contains.
// It removes everything it can but returns the first error
// it encounters.
func (this fileUtil) DeleteAll(filePath string) (bool, error) {
	err := os.RemoveAll(filePath)
	return err == nil, err
}

// CreateNewFile creates the named file with mode 0666 (before umask)
//
// If there is an error, it will be of type *PathError.
func (this fileUtil) CreateNewFile(filePath string, replace... bool) (bool, error) {
	if len(replace) > 0 && replace[0] {
		if this.Exists(filePath) {
			return false, nil
		}
	}
	w, err := os.Create(filePath)
	if err != nil {
		return err == nil, err
	}
	defer w.Close()
	return err == nil, err
}

// RenameTo renames (moves) old path to new path.
// OS-specific restrictions may apply when old path and new path are in different directories.
//
// If there is an error, it will be of type *LinkError.
func (this fileUtil) RenameTo(oldpath string, newpath string, replace... bool) (bool, error) {
	if len(replace) > 0 && replace[0] {
		if this.Exists(newpath) {
			return false, nil
		}
	}
	err := os.Rename(oldpath, newpath)
	return err == nil, err
}

//Parent returns the path string of the specified file path's parent, or empty string if this file path does not name a parent directory
func (this fileUtil) Parent(filePath string) string {
	slashPath := filepath.ToSlash(filePath)
	index := strings.LastIndex(slashPath, slash)
	if index < 0 {
		return ""
	}
	return string(filePath[0: index])
}

// CurrentDirectory returns the current working directory
func (this fileUtil) CurrentDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir, err
}

// UserHome returns the user home directory
func (this fileUtil) UserHome() (string) {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// TempDir returns the default directory to use for temporary files.
func (this fileUtil) TempDir() string {
	return os.TempDir()
}

// FileDir returns the directory that contains the specified file path.
// If the file path is a directory, FileDir returns the path self.
func (this fileUtil) FileDir(filePath string) (string, error) {
	isFile, err := this.IsFile(filePath)
	if err != nil {
		return "", err
	}
	if isFile {
		return this.Parent(filePath), nil
	} else {
		return filePath, nil
	}
}

// IsAbsolute reports whether the path is absolute.
func (this fileUtil) IsAbsolute(filePath string) bool {
	return filepath.IsAbs(filePath)
}

// AbsolutePath returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path.
//
// If there is an error, it will be of type *PathError.
func (this fileUtil) AbsolutePath(filePath string) (string, error) {
	if filepath.IsAbs(filePath) {
		return filepath.Abs(filePath)
	}
	currentDir, err := this.CurrentDirectory()
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(this.JoinPath(currentDir, filePath))
	return abs, err
}

// JoinPath joins any number of path elements into a single path, adding
// a Separator if necessary, all empty strings are ignored.
// On Windows, the result is a UNC path if and only if the first path
// element is a UNC path.
func (this fileUtil) JoinPath(filePath... string) string {
	return filepath.Join(filePath...)
}

func (this fileUtil) CleanPath(filePath string) string {
	return filepath.Clean(filePath)
}

func (this fileUtil) PathEquals(leftPath, rightPath string) bool {
	if this.CleanPath(leftPath) == this.CleanPath(rightPath) {
		return true
	}
	if xi, err := os.Stat(leftPath); err == nil {
		if yi, err := os.Stat(rightPath); err == nil {
			return os.SameFile(xi, yi)
		}
	}
	return false
}

func (this fileUtil)  MD5OfFile(filePath string) []byte {
	fi, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer fi.Close()

	r := bufio.NewReader(fi)

	buf := make([]byte, 1024)
	md5sum := md5.New()
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return nil
		}
		if n == 0 {
			break
		}

		md5sum.Write(buf[:n])
	}

	return md5sum.Sum(nil)
}
