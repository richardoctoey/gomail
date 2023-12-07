package gomail

import "io"

type fileCustom struct {
	Name        string
	Header      map[string][]string
	ObjectBytes []byte
	CopyFunc    func(w io.Writer) error
}

func (f *fileCustom) setHeader(field, value string) {
	f.Header[field] = []string{value}
}
