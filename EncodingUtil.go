package common

import (
	"sync"
	"github.com/axgle/mahonia"
)

type encodingUtil struct {
	mutex sync.Mutex
}

var Encoding = encodingUtil{}

func (this encodingUtil) EncodingToUTF8(content string, encoding string) string {
	return mahonia.NewDecoder(encoding).ConvertString(string(content))
}

func (this encodingUtil) UTF8ToEncoding(content string, encoding string) string {
	return mahonia.NewEncoder(encoding).ConvertString(string(content))
}
